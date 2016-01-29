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
	if p[i].Kind() == p[j].Kind() {
		switch p[i].Kind() {
		case reflect.String:
			return p[i].String() < p[j].String()
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return p[i].Int() < p[j].Int()
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			return p[i].Uint() < p[j].Uint()
		case reflect.UnsafePointer:
			return p[i].UnsafeAddr() < p[j].UnsafeAddr()
		}
	}
	return p[i].Kind() < p[j].Kind()
}
func (p valueSlice) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p valueSlice) Sort() {
	sort.Sort(valueSlice(p))

}
