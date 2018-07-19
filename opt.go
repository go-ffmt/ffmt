package ffmt

import (
	"fmt"
	"io"
	"reflect"
)

type optional struct {
	style style  // Format style
	depth int    // Maximum recursion depth
	opt   option // Option
}

// NewOptional ffmt optional
func NewOptional(depth int, b style, opt option) *optional {
	return &optional{
		style: b,
		opt:   opt,
		depth: depth,
	}
}

func (s *optional) Fprint(w io.Writer, i ...interface{}) (int, error) {
	return fmt.Fprint(w, s.Sprint(i...))
}

func (s *optional) Print(i ...interface{}) (int, error) {
	return fmt.Print(s.Sprint(i...))
}

func (s *optional) Sprint(i ...interface{}) string {
	switch len(i) {
	case 0:
		return ""
	case 1:
		buf := getBuilder()
		defer putBuilder(buf)
		sb := &format{
			buf:      buf,
			filter:   map[uintptr]bool{},
			optional: *s,
		}
		sb.fmt(reflect.ValueOf(i[0]), 0)
		sb.buf.WriteByte('\n')
		ret := sb.buf.String()
		if s.opt.IsCanRowSpan() {
			return Align(ret)
		}
		return ret
	default:
		return s.Sprint(i)
	}
}

type option uint32

// Formatted option
const (
	CanDefaultString   option = 1 << (31 - iota) // can use .(fmt.Stringer)
	CanFilterDuplicate                           // Filter duplicates
	CanRowSpan                                   // Fold line
)

func (t option) IsCanDefaultString() bool {
	return (t & CanDefaultString) != 0
}

func (t option) IsCanFilterDuplicate() bool {
	return (t & CanFilterDuplicate) != 0
}

func (t option) IsCanRowSpan() bool {
	return (t & CanRowSpan) != 0
}

type style int

// Formatted style
const (
	StyleP     style = iota + 1 // Display type and data
	StylePuts                   // Display data
	StylePrint                  // Display data; string without quotes
	StylePjson                  // The json style display; Do not show private
)
