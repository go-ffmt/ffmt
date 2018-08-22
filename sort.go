package ffmt

import (
	"reflect"
	"sort"
)

type valueSlice []reflect.Value

func (p valueSlice) Len() int {
	return len(p)
}
func (p valueSlice) Less(i, j int) bool {
	pi := p[i]
	pj := p[j]
	for pi.Kind() == reflect.Interface || pi.Kind() == reflect.Ptr {
		pi = pi.Elem()
	}
	for pj.Kind() == reflect.Interface || pj.Kind() == reflect.Ptr {
		pj = pj.Elem()
	}
	if pi.Kind() == pj.Kind() {
		switch pi.Kind() {
		case reflect.String:
			return pi.String() < pj.String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return pi.Int() < pj.Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return pi.Uint() < pj.Uint()
		case reflect.Float32, reflect.Float64:
			return pi.Float() < pj.Float()
		default:
			return true
		}
	}
	return pi.Kind() > pj.Kind()
}
func (p valueSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p valueSlice) Sort() {
	sort.Sort(p)
}
