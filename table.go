package ffmt

import (
	"fmt"
	"reflect"
	"strings"
	"unicode"
)

// ToTable Data to table data
func ToTable(t interface{}, is ...interface{}) [][]string {
	r := make([][]string, len(is)+1)
	val := reflect.ValueOf(t)
	val = reflect.Indirect(val)
	typ := val.Type()
	switch val.Kind() {
	case reflect.Struct:
		for i := 0; i != val.NumField(); i++ {
			r[0] = append(r[0], typ.Field(i).Name)
		}
	case reflect.Map:
		ks := val.MapKeys()
		valueSlice(ks).Sort()
		for i := 0; i != len(ks); i++ {
			m := ks[i]
			if m.Kind() == reflect.Ptr || m.Kind() == reflect.Interface {
				m = m.Elem()
			}
			if m.CanInterface() {
				r[0] = append(r[0], fmt.Sprint(m.Interface()))
			}
		}
	default:
		return nil
	}

	for k, v := range is {
		val0 := reflect.ValueOf(v)
		val0 = reflect.Indirect(val0)
		switch val0.Kind() {
		case reflect.Struct:
			for i := 0; i != len(r[0]); i++ {
				field := val0.FieldByName(r[0][i])
				if field.CanInterface() {
					r[k+1] = append(r[k+1], fmt.Sprint(field.Interface()))
				}
			}
		case reflect.Map:
			for i := 0; i != len(r[0]); i++ {
				vv := val0.MapIndex(reflect.ValueOf(r[0][i]))
				if vv.IsValid() {
					r[k+1] = append(r[k+1], fmt.Sprint(vv))
				} else {
					r[k+1] = append(r[k+1], "")
				}
			}
		default:
			return nil
		}
	}
	return r
}

// FmtTable Format table data
func FmtTable(b [][]string) (ss []string) {
	maxs := []int{}
	for _, v1 := range b {
		for k, v2 := range v1 {
			if len(maxs) == k {
				maxs = append(maxs, 0)
			}
			if b := strLen(v2); maxs[k] < b {
				maxs[k] = b
			}
		}
	}
	buf := getBuilder()
	for _, v1 := range b {
		buf.Reset()
		for k, v2 := range v1 {
			buf.WriteString(v2)
			ps := maxs[k] - strLen(v2) + 1
			for i := 0; i != ps; i++ {
				buf.WriteByte(' ')
			}
		}
		ss = append(ss, buf.String())
	}
	putBuilder(buf)
	return
}

// TableText table text
func TableText(b string, prefix, split string) string {
	rows := []string{}
	table := [][]string{}
	for _, v := range strings.Split(b, "\n") {
		if prefix != "" && !strings.HasPrefix(v, prefix) {
			if len(table) != 0 {
				for _, v := range FmtTable(table) {
					rows = append(rows, v)
				}
				table = table[:0]
			}
			rows = append(rows, v)
			continue
		}

		row := []string{}
		ss := strings.Split(v, split)
		for i, col := range ss {
			if i == 0 {
				row = append(row, strings.TrimRightFunc(col, unicode.IsSpace))
			} else {
				row = append(row, strings.TrimFunc(col, unicode.IsSpace))
			}
			if i != len(ss)-1 {
				row[i] = row[i] + split
			}
		}
		table = append(table, row)
	}
	if len(table) != 0 {
		for _, v := range FmtTable(table) {
			rows = append(rows, v)
		}
	}
	ret := strings.Join(rows, "\n")
	return ret
}
