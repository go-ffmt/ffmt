package ffmt

import (
	"fmt"
	"reflect"
	"regexp"
)

func toMap(data interface{}) map[string]string {
	var par = map[string]string{}
	var val = reflect.ValueOf(data)
	val = reflect.Indirect(val)
	switch val.Kind() {
	case reflect.Map:
		for _, k := range val.MapKeys() {
			var v = val.MapIndex(k)
			par[fmt.Sprint(k.Interface())] = fmt.Sprint(v.Interface())
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i != val.Len(); i++ {
			var v = val.Index(i)
			par[fmt.Sprint(i)] = fmt.Sprint(v.Interface())
		}
	case reflect.Struct:
		var typ = val.Type()
		for i := 0; i != typ.NumField(); i++ {
			var f = typ.Field(i)
			var k = f.Name
			var v = val.FieldByName(k)
			if v.CanInterface() {
				par[k] = fmt.Sprint(v.Interface())
			}
		}
	}
	return par
}

func formatMap(str string, par map[string]string) string {
	return regexp.MustCompile(`({[\w\d]+})`).ReplaceAllStringFunc(str, func(s string) string {
		var d, ok = par[s[1:len(s)-1]]
		if ok {
			return d
		}
		return s
	})
}

// Format  Format("hello {name}", "ffmt") to "hello ffmt"
func Format(str string, data ...interface{}) string {
	par := map[string]string{}
	for _, d := range data {
		for k, v := range toMap(d) {
			par[k] = v
		}
	}
	return formatMap(str, par)
}
