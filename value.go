package ffmt

import (
	"fmt"
	"reflect"
	"strings"
)

var DepthMax = 20

const (
	invalid = "<invalid> "
	private = "<private> "
)

func ToString(depth int, b int, i ...interface{}) (ret string) {
	switch len(i) {
	case 0:
		return invalid
	case 1:
		return toString(reflect.ValueOf(i[0]), depth, b)
	default:
		ret += "["
		for k := 0; k != len(i); k++ {
			ret += toString(reflect.ValueOf(i[k]), depth, b)
			ret += " "
		}
		ret += "]"
		return
	}
}

func toString(va reflect.Value, depth int, b int) (ret string) {

	if !va.IsValid() {
		return invalid
	}
	v := va
	if depth <= 0 {
		return toDefault(v, b)
	}

	for v.Kind() == reflect.Ptr {
		if s := getString(va); s != "" {
			return ret + s
		}
		v = v.Elem()
		ret += "&"
	}
	depth--
	switch v.Kind() {
	case reflect.Invalid:
		ret += getName(v, b) + invalid
	case reflect.Struct:
		ret += structToString(v, depth, b)
	case reflect.Map:
		ret += mapToString(v, depth, b)
	case reflect.Array, reflect.Slice:
		ret += sliceToString(v, depth, b)
	case reflect.String:
		ret += stringToString(v, b)
	case reflect.Func:
		ret += funcToString(v)
	case reflect.Chan, reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		ret += pointerToString(v)
	default:
		ret += toDefault(v, b)

	}
	return
}

func toDefault(v reflect.Value, b int) (ret string) {
	if b == 1 {
		ret += getName(v, b)
		ret += "("
	}
	ret += fmt.Sprint(v.Interface())
	if b == 1 {
		ret += ")"
	}
	return
}

func stringToString(v reflect.Value, b int) (ret string) {
	if b == 1 {
		ret += toDefault(v, b)
	} else if b == 2 {
		ret += v.String()
	} else {
		ret += `"`
		ret += strings.Replace(v.String(), `"`, `'`, -1)
		ret += `"`

	}
	return
}

func funcToString(v reflect.Value) (ret string) {
	t := v.Type()
	in := ""
	if t.NumIn() != 0 {
		for i := 0; ; {
			in += t.In(i).String()
			i++
			if i == t.NumIn() {
				break
			}
			in += ","
		}
	}
	out := ""
	if t.NumOut() != 0 {
		for i := 0; ; {
			out += t.Out(i).String()
			i++
			if i == t.NumOut() {
				break
			}
			out += ","
		}
	}

	return fmt.Sprintf("< func(%s)(%s)(0x%020x) > ", in, out, v.Pointer())
}

func pointerToString(v reflect.Value) (ret string) {
	return fmt.Sprintf("%s(0x%020x) ", v.Kind().String(), v.Pointer())
}

func structToString(v reflect.Value, depth int, b int) (ret string) {
	ret += getName(v, b)

	cs := getString(v)
	if cs != "" {
		ret += "< "
		ret += cs
		ret += " >"
	} else {
		ret += "{"
		v.FieldByNameFunc(func(n string) bool {
			ret += n
			ret += ":"
			v0 := v.FieldByName(n)
			if v0.CanInterface() {
				ret += toString(v0, depth, b)
				ret += " "
			} else {
				ret += private
			}
			return false
		})
		ret += "}"
	}

	return
}

func mapToString(v reflect.Value, depth int, b int) (ret string) {
	mk := v.MapKeys()
	ret += getName(v, b)
	ret += "["
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		ret += toString(k, 2, b)
		ret += ":"
		ret += toString(v.MapIndex(k), depth, b)
		ret += " "
	}
	ret += "]"
	return
}

func sliceToString(v reflect.Value, depth int, b int) (ret string) {
	ret += getName(v, b)
	ret += "["
	for i := 0; i != v.Len(); i++ {
		ret += toString(v.Index(i), depth, b)
		ret += " "
	}
	ret += "]"
	return
}

func getString(v reflect.Value) string {
	i := v.Interface()

	if e, b := i.(fmt.Stringer); b {
		return e.String()
	}
	if e, b := i.(fmt.GoStringer); b {
		return e.GoString()
	}
	return ""
}

func getName(v reflect.Value, b int) (ret string) {
	if b != 1 {
		return ""
	}
	t := v.Type()
	if t.PkgPath() != "" {
		ret += t.PkgPath()
		ret += "."
	}

	if t.Name() != "" {
		ret += t.Name()
	} else {
		ret += t.Kind().String()
	}
	return
}
