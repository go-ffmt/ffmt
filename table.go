package ffmt

import (
	"bytes"
	"fmt"
	"reflect"
)

// 制表
func ToTable(t interface{}, is ...interface{}) [][]string {
	r := make([][]string, len(is)+1)
	val := reflect.ValueOf(t)
	typ := val.Type()
	for i := 0; i != val.NumField(); i++ {
		r[0] = append(r[0], typ.Field(i).Name)
	}

	for k, v := range is {
		val := reflect.ValueOf(v)
		for i := 0; i != val.NumField(); i++ {
			r[k+1] = append(r[k+1], fmt.Sprint(val.FieldByName(r[0][i]).Interface()))
		}
	}
	return r
}

// 制表格式化
func FmtTable(b [][]string) (ss []string) {
	maxs := []int{}
	for _, v1 := range b {
		for k, v2 := range v1 {
			if len(maxs) == k {
				maxs = append(maxs, 0)
			}
			if b := Biglen(v2); maxs[k] < b {
				maxs[k] = b
			}
		}
	}
	buf := bytes.NewBuffer(nil)
	for _, v1 := range b {
		buf.Reset()
		for k, v2 := range v1 {
			buf.WriteString(v2)
			ps := maxs[k] - Biglen(v2) + 1
			for i := 0; i != ps; i++ {
				buf.WriteByte(' ')
			}
		}
		ss = append(ss, buf.String())
	}
	return
}
