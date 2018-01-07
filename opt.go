package ffmt

import (
	"bytes"
	"fmt"
	"io"
	"reflect"
)

type optional struct {
	style stlye  // 格式化风格
	depth int    // 最大递归深度
	opt   option // 配置
}

func NewOptional(depth int, b stlye, opt option) *optional {
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
		buf := pool.Get().(*bytes.Buffer)
		buf.Reset()
		defer pool.Put(buf)
		sb := &format{
			buf:      buf,
			filter:   map[uintptr]bool{},
			optional: *s,
		}
		sb.fmt(reflect.ValueOf(i[0]), 0)
		sb.buf.WriteByte('\n')
		ret := sb.buf.String()
		if s.opt.IsCanRowSpan() {
			return nodes(ret)
		}
		return ret
	default:
		return s.Sprint(i)
	}
}

type option uint32

const (
	CanDefaultString option = 1 << (31 - iota)
	CanFilterDuplicate
	CanRowSpan
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

type stlye int

const (
	StlyeP     stlye = iota + 1 // 显示完整的类型名和数据
	StlyePuts                   // 显示数据
	StlyePrint                  // 显示数据 字符串不加引号
	StlyePjson                  // 以json风格显示数据 不显示结构体私有项
)
