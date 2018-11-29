package ffmt

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"unicode"
	"unicode/utf8"
)

const (
	invalidJSON = "null"
	invalid     = "<nil>"
	private     = "<private>"
)

type format struct {
	optional
	buf    builder
	filter map[uintptr]bool
}

func (s *format) fmt(va reflect.Value, depth int) {
	if !va.IsValid() {
		s.nilBuf()
		return
	}

	if va.Kind() == reflect.Ptr && va.IsNil() {
		s.nilBuf()
		return
	}

	v := va
	if depth >= s.depth {
		s.defaultBuf(v)
		return
	}

	if s.getString(v) {
		return
	}

	for v.Kind() == reflect.Ptr {
		if s.opt.IsCanDefaultString() {
			if s.filter[v.Pointer()] {
				s.xxBuf(v, v.Pointer())
				return
			}
			s.filter[v.Pointer()] = true
		}

		v = v.Elem()
		if !v.IsValid() {
			s.nilBuf()
			return
		}
		switch s.style {
		case StylePjson:
		default:
			s.buf.WriteByte('&')
		}
		if s.getString(va) {
			return
		}
	}
	s.switchType(v, depth)
	return
}

func (s *format) switchType(v reflect.Value, depth int) {
	switch v.Kind() {
	case reflect.Invalid:
		s.nilBuf()
	case reflect.Struct:
		switch s.style {
		case StylePjson:
			s.mapBuf(reflect.ValueOf(struct2Map(v)), depth)
		default:
			s.structBuf(v, depth)
		}
	case reflect.Map:
		s.mapBuf(v, depth)
	case reflect.Array, reflect.Slice:
		s.sliceBuf(v, depth)
	case reflect.String:
		s.stringBuf(v)
	case reflect.Func:
		s.funcBuf(v)
	case reflect.Uintptr:
		s.xxBuf(v, v.Uint())
	case reflect.Chan, reflect.Ptr, reflect.UnsafePointer:
		s.xxBuf(v, v.Pointer())
	case reflect.Interface:
		v = v.Elem()
		if v.IsValid() {
			s.fmt(v, depth)
		} else {
			s.nilBuf()
		}
	default:
		s.defaultBuf(v)
	}
	return
}

// depthBuf write buffer with depth
func (s *format) depthBuf(i int) {
	s.buf.WriteByte('\n')
	for k := 0; k < i; k++ {
		s.buf.WriteByte(Space)
	}
	return
}

// nilBuf write buffer with nil
func (s *format) nilBuf() {
	switch s.style {
	case StylePjson:
		s.buf.WriteString(invalidJSON)
	default:
		s.buf.WriteString(invalid)
	}
	return
}

// defaultBuf write buffer with default string
func (s *format) defaultBuf(v reflect.Value) {
	s.nameBuf(v.Type())
	switch s.style {
	case StyleP:
		s.buf.WriteByte('(')
		fmt.Fprintf(s.buf, "%#v", v.Interface())
		s.buf.WriteByte(')')
	case StylePjson:
		d, _ := json.Marshal(v.Interface())
		s.buf.Write(d)
	default:
		fmt.Fprintf(s.buf, "%+v", v.Interface())
	}
	return
}

// stringBuf write buffer with string
func (s *format) stringBuf(v reflect.Value) {
	switch s.style {
	case StyleP:
		s.defaultBuf(v)
	case StylePuts, StylePjson:
		s.buf.WriteString(strconv.Quote(v.String()))
	default:
		s.buf.WriteString(v.String())
	}
	return
}

// funcBuf write buffer with func address
func (s *format) funcBuf(v reflect.Value) {
	switch s.style {
	case StylePjson:
		s.buf.WriteByte('"')
		defer s.buf.WriteByte('"')
	case StyleP:
		s.buf.WriteByte('<')
		defer s.buf.WriteByte('>')
	}
	s.buf.WriteString("func(")
	t := v.Type()
	if t.NumIn() != 0 {
		for i := 0; ; {
			s.buf.WriteString(t.In(i).String())
			i++
			if i == t.NumIn() {
				break
			}
			s.buf.WriteByte(',')
		}
	}
	s.buf.WriteString(")(")
	if t.NumOut() != 0 {
		for i := 0; ; {
			s.buf.WriteString(t.Out(i).String())
			i++
			if i == t.NumOut() {
				break
			}
			s.buf.WriteByte(',')
		}
	}
	s.buf.WriteByte(')')
	fmt.Fprintf(s.buf, "(0x%020x)", v.Pointer())
	return
}

// xxBuf write buffer with hex format
func (s *format) xxBuf(v reflect.Value, i interface{}) {
	switch s.style {
	case StylePjson:
		s.buf.WriteByte('"')
		defer s.buf.WriteByte('"')
	case StyleP:
		s.buf.WriteByte('<')
		defer s.buf.WriteByte('>')
	}
	s.nameBuf(v.Type())
	fmt.Fprintf(s.buf, "(0x%020x)", i)
	return
}

// structBuf write buffer with struct
func (s *format) structBuf(v reflect.Value, depth int) {
	s.nameBuf(v.Type())
	s.buf.WriteByte('{')
	t := v.Type()
	for i := 0; i != t.NumField(); i++ {
		f := t.Field(i)
		v0 := v.Field(i)
		s.depthBuf(depth + 1)
		s.buf.WriteString(f.Name)
		s.buf.WriteString(colSym)
		if isPrivateName(f.Name) {
			s.buf.WriteString(private)
		} else {
			s.fmt(v0, depth+1)
		}
	}
	s.depthBuf(depth)
	s.buf.WriteByte('}')
	return
}

// mapBuf write buffer with map
func (s *format) mapBuf(v reflect.Value, depth int) {
	mk := v.MapKeys()
	valueSlice(mk).Sort()
	s.nameBuf(v.Type())
	s.buf.WriteByte('{')
	for i := 0; i != len(mk); i++ {
		k := mk[i]
		switch s.style {
		case StylePjson:
			if i != 0 {
				s.depthBuf(depth)
				s.buf.WriteByte(',')
			} else {
				s.depthBuf(depth + 1)
			}
		default:
			s.depthBuf(depth + 1)
		}
		s.fmt(k, s.depth-1)
		s.buf.WriteString(colSym)
		s.fmt(v.MapIndex(k), depth+1)
	}
	s.depthBuf(depth)
	s.buf.WriteByte('}')
	return
}

// sliceBuf write buffer with slice
func (s *format) sliceBuf(v reflect.Value, depth int) {
	s.nameBuf(v.Type())
	s.buf.WriteByte('[')
	for i := 0; i != v.Len(); i++ {
		switch s.style {
		case StylePjson:
			if i != 0 {
				s.depthBuf(depth)
				s.buf.WriteByte(',')
			} else {
				s.depthBuf(depth + 1)
			}
		default:
			s.depthBuf(depth + 1)
		}
		s.fmt(v.Index(i), depth+1)
	}
	s.depthBuf(depth)
	s.buf.WriteByte(']')
	return
}

// nameBuf write buffer with type name
func (s *format) nameBuf(t reflect.Type) bool {
	switch s.style {
	case StyleP:
		switch t.Kind() {
		case reflect.Map:
			s.buf.WriteString("map[")
			s.nameBuf(t.Key())
			s.buf.WriteByte(']')
			return s.nameBuf(t.Elem())
		case reflect.Slice:
			s.buf.WriteString("[]")
			return s.nameBuf(t.Elem())
		case reflect.Array:
			s.buf.WriteByte('[')
			s.buf.WriteString(strconv.FormatInt(int64(t.Len()), 10))
			s.buf.WriteByte(']')
			return s.nameBuf(t.Elem())
		case reflect.Ptr:
			s.buf.WriteByte('*')
			return s.nameBuf(t.Elem())
		default:
			if pkg := t.PkgPath(); pkg != "" {
				s.buf.WriteString(pkg)
				s.buf.WriteByte('.')
			}

			if t.Name() != "" {
				s.buf.WriteString(t.Name())
			} else {
				s.buf.WriteString(t.String())
			}
			return true
		}
	}
	return false
}

// getString write buffer with default string
// returns true if can't default string
func (s *format) getString(v reflect.Value) bool {
	if !s.opt.IsCanDefaultString() {
		return false
	}

	switch s.style {
	case StylePjson:
		r := getString(v)
		if r == "" {
			return false
		}
		vv, _ := json.Marshal(v.Interface())
		s.buf.Write(vv)
	case StyleP:
		r := getString(v)
		if r == "" {
			return false
		}
		s.buf.WriteByte('<')
		s.buf.WriteString(r)
		s.buf.WriteByte('>')
	default:
		r := getString(v)
		if r == "" {
			return false
		}
		s.buf.WriteString(r)
	}
	return true
}

// getString return default string
func getString(v reflect.Value) string {

	if v.Kind() == reflect.Interface {
		if v.IsNil() {
			return ""
		}
		return getString(v.Elem())
	}

	if !v.CanInterface() {
		return ""
	}

	i := v.Interface()

	if e, b := i.(fmt.Stringer); b && e != nil {
		return e.String()
	}
	if e, b := i.(fmt.GoStringer); b && e != nil {
		return e.GoString()
	}
	return ""
}

// isPrivateName returns it is a private name
func isPrivateName(name string) bool {
	ch, _ := utf8.DecodeRuneInString(name)
	return !unicode.IsUpper(ch)
}

// struct2Map returns map from struct
func struct2Map(v reflect.Value) map[string]interface{} {
	t := v.Type()
	data := map[string]interface{}{}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if isPrivateName(f.Name) {
			continue
		}
		data[f.Name] = v.Field(i).Interface()
	}
	return data
}
