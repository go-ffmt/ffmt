package ffmt

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
)

var DepthMax = 20

const (
	invalid = "<nil> "
	private = "<private> "
)

type stlye int

const (
	sp stlye = iota + 1
	sputs
	sprint
)

func toString(depth int, b stlye, i ...interface{}) (ret string) {
	sb := &sbuf{style: b}
	switch len(i) {
	case 0:
		return
	case 1:
		sb.fmt(reflect.ValueOf(i[0]), depth)
		return sb.String()
	default:
		sb.WriteByte('[')
		for k := 0; k != len(i); k++ {
			sb.fmt(reflect.ValueOf(i[k]), depth)
			sb.WriteByte(' ')
		}
		sb.WriteByte(']')
		return sb.String()
	}
	return
}

type sbuf struct {
	bytes.Buffer
	style stlye
}

func (s *sbuf) fmt(va reflect.Value, depth int) {
	if !va.IsValid() {
		s.WriteString(invalid)
		return
	}
	v := va
	if depth <= 0 {
		s.toDefault(v)
		return
	}
	if s.getString(va) {
		return
	}
	for v.Kind() == reflect.Ptr {
		v = v.Elem()
		s.WriteByte('&')
		if s.getString(va) {
			return
		}
	}
	depth--
	switch v.Kind() {
	case reflect.Invalid:
		s.WriteString(invalid)
	case reflect.Struct:
		s.struct2String(v, depth)
	case reflect.Map:
		s.map2String(v, depth)
	case reflect.Array, reflect.Slice:
		s.slice2String(v, depth)
	case reflect.String:
		s.string2String(v)
	case reflect.Func:
		s.func2String(v)
	case reflect.Chan, reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		s.pointer2String(v)
	default:
		s.toDefault(v)
	}
	return
}

func (s *sbuf) toDefault(v reflect.Value) {
	switch s.style {
	case sp:
		s.getName(v)
		s.WriteByte('(')
		s.WriteString(fmt.Sprint(v.Interface()))
		s.WriteByte(')')
	default:
		s.WriteString(fmt.Sprint(v.Interface()))
	}
	return
}

func (s *sbuf) string2String(v reflect.Value) {
	switch s.style {
	case sp:
		s.toDefault(v)
	case sputs:
		s.WriteByte('"')
		s.WriteString(strings.Replace(v.String(), `"`, `'`, -1))
		s.WriteByte('"')
	default:
		s.WriteString(v.String())
	}
	return
}

func (s *sbuf) func2String(v reflect.Value) {
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

	s.WriteString(fmt.Sprintf("<func(%s)(%s)(0x%020x)> ", in, out, v.Pointer()))
	return
}

func (s *sbuf) pointer2String(v reflect.Value) {
	s.WriteString(fmt.Sprintf("%s(0x%020x) ", v.Kind().String(), v.Pointer()))
	return
}

func (s *sbuf) struct2String(v reflect.Value, depth int) {
	s.getName(v)
	s.WriteByte('{')
	t := v.Type()
	for i := 0; i != t.NumField(); i++ {
		f := t.Field(i)
		s.WriteString(f.Name)
		s.WriteByte(':')
		v0 := v.Field(i)
		if v0.CanInterface() {
			s.fmt(v0, depth)
			s.WriteByte(' ')
		} else {
			s.WriteString(private)
		}
	}
	s.WriteByte('}')
	return
}

func (s *sbuf) map2String(v reflect.Value, depth int) {
	mk := v.MapKeys()
	s.getName(v)
	s.WriteByte('[')
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		s.fmt(k, 2)
		s.WriteByte(':')
		s.fmt(v.MapIndex(k), depth)
		s.WriteByte(' ')
	}
	s.WriteByte(']')
	return
}

func (s *sbuf) slice2String(v reflect.Value, depth int) {
	s.getName(v)
	s.WriteByte('[')
	for i := 0; i != v.Len(); i++ {
		s.fmt(v.Index(i), depth)
		s.WriteByte(' ')
	}
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
		s.WriteByte('<')
		s.WriteString(r)
		s.WriteByte('>')
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
