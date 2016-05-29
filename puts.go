package ffmt

import (
	"fmt"
	"io"
)

const colSym = ": "

var DepthMax = 5
var Space byte = ' '

// P 系列 更改风格 显示完整的类型
func Fp(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sp(a...))
}

func P(a ...interface{}) (int, error) {
	return fmt.Print(Sp(a...))
}

func Sp(a ...interface{}) string {
	return nodes(toString(DepthMax, sp, a...))
}

// Puts 系列 更改风格
func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sputs(a...))
}

func Puts(a ...interface{}) (int, error) {
	return fmt.Print(Sputs(a...))
}

func Sputs(a ...interface{}) string {
	return nodes(toString(DepthMax, sputs, a...))
}

// Print 系列 默认风格
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprint(a...))
}

func Print(a ...interface{}) (int, error) {
	return fmt.Print(Sprint(a...))
}

func Sprint(a ...interface{}) string {
	return nodes(toString(DepthMax, sprint, a...))
}

// Pjson 系列
func Fpjson(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Spjson(a...))
}

func Pjson(a ...interface{}) (int, error) {
	return fmt.Print(Spjson(a...))
}

func Spjson(a ...interface{}) string {
	return nodes(toString(DepthMax, spjson, a...))
}
