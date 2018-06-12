package ffmt

import (
	"io"
)

const colSym = ": "

// Space rune
var Space byte = ' '

// p go style display types for debug
var defD = NewOptional(10, StyleP, CanFilterDuplicate|CanRowSpan)

// D for debug
func D(a ...interface{}) {
	MarkStack(1, defD.Sprint(a...))
}

// Sd for debug
func Sd(a ...interface{}) string {
	return SmarkStack(1, defD.Sprint(a...))
}

// P go style display types
var defP = NewOptional(5, StyleP, CanDefaultString|CanFilterDuplicate|CanRowSpan)

// Fp The go style friendly display types and data to writer
func Fp(w io.Writer, a ...interface{}) (int, error) {
	return defP.Fprint(w, a...)
}

// P The go style friendly display types and data
func P(a ...interface{}) (int, error) {
	return defP.Print(a...)
}

// Sp The go style friendly display types and data to string
func Sp(a ...interface{}) string {
	return defP.Sprint(a...)
}

// Puts go style
var defPuts = NewOptional(5, StylePuts, CanDefaultString|CanFilterDuplicate|CanRowSpan)

// Fputs The go style friendly to writer
func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return defPuts.Fprint(w, a...)
}

// Puts The go style friendly display
func Puts(a ...interface{}) (int, error) {
	return defPuts.Print(a...)
}

// Sputs The go style friendly to string
func Sputs(a ...interface{}) string {
	return defPuts.Sprint(a...)
}

// Print go style
var defPrint = NewOptional(5, StylePrint, CanDefaultString|CanFilterDuplicate|CanRowSpan)

// Fprint The go style friendly to writer
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return defPrint.Fprint(w, a...)
}

// Print The go style friendly display
func Print(a ...interface{}) (int, error) {
	return defPrint.Print(a...)
}

// Sprint The go style friendly to string
func Sprint(a ...interface{}) string {
	return defPrint.Sprint(a...)
}

// Pjson json style
var defPjson = NewOptional(20, StylePjson, CanDefaultString|CanRowSpan)

// Fpjson The json style friendly display to writer
func Fpjson(w io.Writer, a ...interface{}) (int, error) {
	return defPjson.Fprint(w, a...)
}

// Pjson The json style friendly display
func Pjson(a ...interface{}) (int, error) {
	return defPjson.Print(a...)
}

// Spjson The json style friendly display to string
func Spjson(a ...interface{}) string {
	return defPjson.Sprint(a...)
}
