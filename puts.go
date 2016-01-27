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
	return stringToNode(fmts(toString(DepthMax, sputs, a...), 0, 0)).String()
}

// P 系列 更改风格 显示完整的类型
func Fp(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sp(a...))
}

func P(a ...interface{}) (int, error) {
	return fmt.Print(Sp(a...))
}

func Sp(a ...interface{}) string {
	return stringToNode(fmts(toString(DepthMax, sp, a...), 0, 0)).String()
}

// Print 系列 默认风格 显示有点糟糕 用来格式化json 还是不错的
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sprint(a...))
}

func Print(a ...interface{}) (int, error) {
	return fmt.Print(Sprint(a...))
}

func Sprint(a ...interface{}) string {
	return stringToNode(fmts(toString(DepthMax, sprint, a...), 0, 0)).String()
}
