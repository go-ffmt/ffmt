package ffmt

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

var DepthMax = 5

const (
	null    = "null"
	invalid = "<nil>"
	private = "<private>"
)

type stlye int

const (
	sp stlye = iota + 1
	sputs
	sprint
	sjson
)

func toString(depth int, b stlye, i ...interface{}) string {
	switch len(i) {
	case 0:
		return ""
	case 1:
		sb := &sbuf{style: b, depth: depth, buf: map[uintptr]bool{}}
		sb.fmt(reflect.ValueOf(i[0]), 0)
		sb.WriteByte('\n')
		return sb.String()
	default:
		return toString(depth, b, i)
	}
}

type sbuf struct {
	bytes.Buffer
	style stlye
	depth int
	buf   map[uintptr]bool
}

func (s *sbuf) fmt(va reflect.Value, depth int) {
	if !va.IsValid() {
		switch s.style {
		case sjson:
			s.WriteString(null)
		default:
			s.WriteString(invalid)
		}
		return
	}
	v := va
	if depth >= DepthMax {
		s.toDefault(v)
		return
	}

	if s.getString(v) {
		return
	}

	for v.Kind() == reflect.Ptr {
		if s.buf[v.Pointer()] {
			s.pointer2String(v)
			return
		} else {
			s.buf[v.Pointer()] = true
		}

		v = v.Elem()
		if !v.IsValid() {
			switch s.style {
			case sjson:
				s.WriteString(null)
			default:
				s.WriteString(invalid)
			}
			return
		}
		switch s.style {
		case sjson:
		default:
			s.WriteByte('&')
		}
		if s.getString(va) {
			return
		}
	}
	switch v.Kind() {
	case reflect.Invalid:
		s.WriteString(invalid)
	case reflect.Struct:
		switch s.style {
		case sjson:
			s.map2String(reflect.ValueOf(struct2Map(v)), depth)
		default:
			s.struct2String(v, depth)
		}
	case reflect.Map:
		s.map2String(v, depth)
	case reflect.Array, reflect.Slice:
		s.slice2String(v, depth)
	case reflect.String:
		s.string2String(v)
	case reflect.Func:
		s.func2String(v)
	case reflect.Uintptr, reflect.UnsafePointer:
		s.x2string(v)
	case reflect.Chan, reflect.Ptr:
		s.pointer2String(v)
	case reflect.Interface:
		v = v.Elem()
		if v.IsValid() {
			s.fmt(v, depth)
		} else {
			s.WriteString(null)
		}
	default:
		s.toDefault(v)
	}
	return
}

func (s *sbuf) toDepth(i int) {
	s.WriteByte('\n')
	s.getSpace(i)
}
func (s *sbuf) getSpace(i int) {
	for k := 0; k < i; k++ {
		s.WriteByte(' ')
	}
}

func (s *sbuf) toDefault(v reflect.Value) {
	switch s.style {
	case sp:
		s.getName(v)
		s.WriteByte('(')
		s.WriteString(fmt.Sprint(v.Interface()))
		s.WriteByte(')')
	case sjson:
		js, _ := json.Marshal(v.Interface())
		s.WriteString(string(js))
	default:
		s.WriteString(fmt.Sprint(v.Interface()))
	}
	return
}

func (s *sbuf) string2String(v reflect.Value) {
	switch s.style {
	case sp:
		s.toDefault(v)
	case sputs, sjson:
		s.WriteByte('"')
		s.WriteString(strings.Replace(v.String(), `"`, `'`, -1))
		s.WriteByte('"')
	default:
		s.WriteString(v.String())
	}
	return
}

func (s *sbuf) func2String(v reflect.Value) {
	switch s.style {
	case sjson:
		s.WriteByte('"')
		defer s.WriteByte('"')
	default:
		s.WriteByte('<')
		defer s.WriteByte('>')
	}
	s.WriteString("func(")
	t := v.Type()
	if t.NumIn() != 0 {
		for i := 0; ; {
			s.WriteString(t.In(i).String())
			i++
			if i == t.NumIn() {
				break
			}
			s.WriteByte(',')
		}
	}
	s.WriteByte(')')
	s.WriteByte('(')
	if t.NumOut() != 0 {
		for i := 0; ; {
			s.WriteString(t.Out(i).String())
			i++
			if i == t.NumOut() {
				break
			}
			s.WriteByte(',')
		}
	}
	s.WriteByte(')')
	s.WriteString(fmt.Sprintf("(0x%020x)", v.Pointer()))
	return
}

func (s *sbuf) x2string(v reflect.Value) {
	switch s.style {
	case sjson:
		s.WriteByte('"')
		defer s.WriteByte('"')

	default:
		s.WriteByte('<')
		defer s.WriteByte('>')
	}
	s.WriteString(v.Kind().String())
	s.WriteString(fmt.Sprintf("(0x%020x)", v.Uint()))
	return
}

func (s *sbuf) pointer2String(v reflect.Value) {
	switch s.style {
	case sjson:
		s.WriteByte('"')
		defer s.WriteByte('"')

	default:
		s.WriteByte('<')
		defer s.WriteByte('>')
	}
	s.WriteString(v.Kind().String())
	s.WriteString(fmt.Sprintf("(0x%020x)", v.Pointer()))
	return
}

func (s *sbuf) struct2String(v reflect.Value, depth int) {
	s.getName(v)
	s.WriteByte('{')
	t := v.Type()

	for i := 0; i != t.NumField(); i++ {
		f := t.Field(i)
		n := f.Name[0]
		s.toDepth(depth + 1)
		s.WriteString(f.Name)
		s.WriteByte(':')
		s.WriteByte(' ')
		v0 := v.Field(i)
		if n < 'A' || n > 'Z' {
			s.WriteString(private)
		} else {
			s.fmt(v0, depth+1)
		}
	}
	s.toDepth(depth)
	s.WriteByte('}')
	return
}

func (s *sbuf) map2String(v reflect.Value, depth int) {
	mk := v.MapKeys()
	valueSlice(mk).Sort()

	s.getName(v)
	s.WriteByte('{')
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		switch s.style {
		case sjson:
			if i != 0 {
				s.toDepth(depth)
				s.WriteByte(',')
			} else {
				s.toDepth(depth + 1)
			}
		default:
			s.toDepth(depth + 1)
		}
		s.fmt(k, DepthMax-1)
		s.WriteByte(':')
		s.WriteByte(' ')
		s.fmt(v.MapIndex(k), depth+1)
	}
	s.toDepth(depth)
	s.WriteByte('}')
	return
}

func (s *sbuf) slice2String(v reflect.Value, depth int) {
	s.getName(v)
	s.WriteByte('[')
	for i := 0; i != v.Len(); i++ {
		switch s.style {
		case sjson:
			if i != 0 {
				s.toDepth(depth)
				s.WriteByte(',')
			} else {
				s.toDepth(depth + 1)
			}
		default:
			s.toDepth(depth + 1)
		}
		s.fmt(v.Index(i), depth+1)
	}
	s.toDepth(depth)
	s.WriteByte(']')
	return
}

func (s *sbuf) getName(v reflect.Value) {
	switch s.style {
	case sp:
		t := v.Type()
		if t.PkgPath() != "" {
			s.WriteString(t.PkgPath())
			s.WriteByte('.')
		}

		if t.Name() != "" {
			s.WriteString(t.Name())
		} else {
			s.WriteString(t.Kind().String())
		}
	}
	return
}
func (s *sbuf) getString(v reflect.Value) bool {
	if r := getString(v); r != "" {
		switch s.style {
		case sjson:
			s.WriteByte('"')
			s.WriteString(r)
			s.WriteByte('"')
		default:
			s.WriteByte('<')
			s.WriteString(r)
			s.WriteByte('>')
		}
		return true
	}
	return false
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

func struct2Map(v reflect.Value) map[string]interface{} {
	t := v.Type()
	data := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		n := f.Name[0]
		if n < 'A' || n > 'Z' {
			continue
		}
		data[f.Name] = v.Field(i).Interface()
	}
	return data
}
