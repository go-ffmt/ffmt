package ffmt

import "io"

const colSym = ": "

var Space byte = ' '

// P 系列 更改风格 显示完整的类型
var defP = NewOptional(5, StlyeP, CanDefaultString|CanFilterDuplicate|CanRowSpan)

func Fp(w io.Writer, a ...interface{}) (int, error) {
	return defP.Fprint(w, a...)
}

func P(a ...interface{}) (int, error) {
	return defP.Print(a...)
}

func Sp(a ...interface{}) string {
	return defP.Sprint(a...)
}

// Puts 系列 更改风格
var defPuts = NewOptional(5, StlyePuts, CanDefaultString|CanFilterDuplicate|CanRowSpan)

func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return defPuts.Fprint(w, a...)
}

func Puts(a ...interface{}) (int, error) {
	return defPuts.Print(a...)
}

func Sputs(a ...interface{}) string {
	return defPuts.Sprint(a...)
}

// Print 系列 默认风格
var defPrint = NewOptional(5, StlyePrint, CanDefaultString|CanFilterDuplicate|CanRowSpan)

func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return defPrint.Fprint(w, a...)
}

func Print(a ...interface{}) (int, error) {
	return defPrint.Print(a...)
}

func Sprint(a ...interface{}) string {
	return defPrint.Sprint(a...)
}

// Pjson 系列
var defPjson = NewOptional(20, StlyePjson, CanDefaultString|CanRowSpan)

func Fpjson(w io.Writer, a ...interface{}) (int, error) {
	return defPjson.Fprint(w, a...)
}

func Pjson(a ...interface{}) (int, error) {
	return defPjson.Print(a...)
}

func Spjson(a ...interface{}) string {
	return defPjson.Sprint(a...)
}
