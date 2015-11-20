package ffmt

import (
	"fmt"
	"reflect"
)

func ToString(i ...interface{}) (ret string) {
	switch len(i) {
	case 0:
		return "<none>"
	case 1:
		return toString(reflect.ValueOf(i[0]))
	default:
		ret += "["
		for k := 0; k != len(i); k++ {
			ret += toString(reflect.ValueOf(i[k]))
			ret += " "
		}
		ret += "]"
		return
	}
}

func toString(va reflect.Value) (ret string) {
	if !va.IsValid() {
		return "<invalid>"
	}
	v := va

	for v.Kind() == reflect.Ptr {
		v = v.Elem()
		ret += "&"
	}
	switch v.Kind() {
	case reflect.Struct:
		ret += structToString(v)
	case reflect.Map:
		ret += mapToString(v)
	case reflect.Array, reflect.Slice:
		ret += sliceToString(v)
	case reflect.String:
		ret += "\"" + v.String() + "\""
	case reflect.Func:
		ret += funcToString(v)
	case reflect.Chan, reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		ret += pointerToString(v)
	default:
		ret += v.Kind().String() + "(" + fmt.Sprint(v.Interface()) + ")"
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

	return fmt.Sprintf("func(%s)(%s)(0x%020x) ", in, out, v.Pointer())
}

func pointerToString(v reflect.Value) (ret string) {
	return fmt.Sprintf("%s(0x%020x) ", v.Kind().String(), v.Pointer())
}

func structToString(v reflect.Value) (ret string) {
	t := v.Type()
	if t.PkgPath() != "" {
		ret += t.PkgPath()
	} else {
		ret += "main"
	}
	ret += "."
	if t.Name() != "" {
		ret += t.Name()
	} else {
		ret += "<anonym>"
	}
	ret += "{ "
	cs := callString(v)
	if cs != "" {
		ret += cs
	} else {
		v.FieldByNameFunc(func(n string) bool {
			ret += n
			ret += ":"
			v0 := v.FieldByName(n)
			if v0.CanInterface() {
				ret += toString(v0)
				ret += " "
			} else {
				ret += "<private> "
			}
			return false
		})
	}
	ret += "}"
	return
}

func mapToString(v reflect.Value) (ret string) {
	mk := v.MapKeys()
	ret += "map[ "
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		ret += toString(k)
		ret += ":"
		ret += toString(v.MapIndex(k))
		ret += " "
	}
	ret += "] "
	return
}

func sliceToString(v reflect.Value) (ret string) {
	ret += v.Kind().String()
	ret += "[ "
	for i := 0; i != v.Len(); i++ {
		ret += toString(v.Index(i))
		ret += " "
	}
	ret += "] "
	return
}

func callString(v reflect.Value) string {
	i := v.Interface()

	if e, b := i.(fmt.Stringer); b {
		return e.String()
	}
	if e, b := i.(fmt.GoStringer); b {
		return e.GoString()
	}
	return ""
}
