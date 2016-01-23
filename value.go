package ffmt

import (
	"bytes"
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

	sb := &sbuf{}
	switch len(i) {
	case 0:
		return invalid
	case 1:
		sb.toString(reflect.ValueOf(i[0]), depth, b)
		return sb.String()
	default:
		ret += "["
		for k := 0; k != len(i); k++ {
			sb.toString(reflect.ValueOf(i[k]), depth, b)
			ret += sb.String()
			ret += " "
		}
		ret += "]"
		return
	}
}

type sbuf struct {
	bytes.Buffer
}

func (s *sbuf) toString(va reflect.Value, depth int, b int) {

	if !va.IsValid() {
		s.WriteString(invalid)
		return
	}
	v := va
	if depth <= 0 {
		s.toDefault(v, b)
		return
	}

	for v.Kind() == reflect.Ptr {
		if r := getString(va); r != "" {
			s.WriteString(r)
			return
		}
		v = v.Elem()
		s.WriteByte('&')
	}
	depth--
	switch v.Kind() {
	case reflect.Invalid:
		s.getName(v, b)
		s.WriteString(invalid)
	case reflect.Struct:
		s.structToString(v, depth, b)
	case reflect.Map:
		s.mapToString(v, depth, b)
	case reflect.Array, reflect.Slice:
		s.sliceToString(v, depth, b)
	case reflect.String:
		s.stringToString(v, b)
	case reflect.Func:
		s.funcToString(v)
	case reflect.Chan, reflect.Uintptr, reflect.Ptr, reflect.UnsafePointer:
		s.pointerToString(v)
	default:
		s.toDefault(v, b)

	}
	return
}

func (s *sbuf) toDefault(v reflect.Value, b int) {
	if b == 1 {
		s.getName(v, b)
		s.WriteByte('(')
	}
	s.WriteString(fmt.Sprint(v.Interface()))
	if b == 1 {
		s.WriteByte(')')
	}
	return
}

func (s *sbuf) stringToString(v reflect.Value, b int) {
	if b == 1 {
		s.toDefault(v, b)
	} else if b == 2 {
		s.WriteString(v.String())
	} else {
		s.WriteByte('"')
		s.WriteString(strings.Replace(v.String(), `"`, `'`, -1))
		s.WriteByte('"')

	}
	return
}

func (s *sbuf) funcToString(v reflect.Value) {
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

	s.WriteString(fmt.Sprintf("< func(%s)(%s)(0x%020x) > ", in, out, v.Pointer()))
	return
}

func (s *sbuf) pointerToString(v reflect.Value) {
	s.WriteString(fmt.Sprintf("%s(0x%020x) ", v.Kind().String(), v.Pointer()))
	return
}

func (s *sbuf) structToString(v reflect.Value, depth int, b int) {
	s.getName(v, b)
	cs := getString(v)
	if cs != "" {
		s.WriteByte('<')
		s.WriteString(cs)
		s.WriteByte('>')
	} else {
		s.WriteByte('{')
		t := v.Type()
		for i := 0; i != t.NumField(); i++ {
			f := t.Field(i)
			s.WriteString(f.Name)
			s.WriteByte(':')
			v0 := v.Field(i)
			if v0.CanInterface() {
				s.toString(v0, depth, b)
				s.WriteByte(' ')
			} else {
				s.WriteString(private)
			}
		}
		s.WriteByte('}')
	}
	return
}

func (s *sbuf) mapToString(v reflect.Value, depth int, b int) {
	mk := v.MapKeys()
	s.getName(v, b)
	s.WriteByte('[')
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		s.toString(k, 2, b)
		s.WriteByte(':')
		s.toString(v.MapIndex(k), depth, b)
		s.WriteByte(' ')
	}
	s.WriteByte(']')
	return
}

func (s *sbuf) sliceToString(v reflect.Value, depth int, b int) {
	s.getName(v, b)
	s.WriteByte('[')
	for i := 0; i != v.Len(); i++ {
		s.toString(v.Index(i), depth, b)
		s.WriteByte(' ')
	}
	s.WriteByte(']')
	return
}

func (s *sbuf) getName(v reflect.Value, b int) {
	if b != 1 {
		return
	}
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
