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
	default:
		ret += fmt.Sprintln(v.Interface())
	}

	return
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
	f := v.MethodByName("String")

	if f.Kind() == reflect.Func && f.Type().NumIn() == 0 && f.Type().NumOut() == 1 && f.Type().Out(0).Kind() == reflect.String {
		return f.Call([]reflect.Value{})[0].String()
	}
	return ""
}