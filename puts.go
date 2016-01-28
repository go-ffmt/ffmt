package ffmt

import (
	"fmt"
	"io"
)

// Puts 系列 更改风格
func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sputs(a...))
}

func Puts(a ...interface{}) (int, error) {
	return fmt.Print(Sputs(a...))
}

func Sputs(a ...interface{}) string {
	return stringToNode(toString(DepthMax, sputs, a...)).String()
}

// P 系列 更改风格 显示完整的类型
func Fp(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sp(a...))
}

func P(a ...interface{}) (int, error) {
	return fmt.Print(Sp(a...))
}

func Sp(a ...interface{}) string {
	return stringToNode(toString(DepthMax, sp, a...)).String()
}

// Print 系列 默认风格
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprint(a...))
}

func Print(a ...interface{}) (int, error) {
	return fmt.Print(Sprint(a...))
}

func Sprint(a ...interface{}) string {
	return stringToNode(toString(DepthMax, sprint, a...)).String()
}

// Json 系列
func Fjson(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sjson(a...))
}

func Json(a ...interface{}) (int, error) {
	return fmt.Print(Sjson(a...))
}

func Sjson(a ...interface{}) string {
	return stringToNode(toString(DepthMax, sjson, a...)).String()
}
