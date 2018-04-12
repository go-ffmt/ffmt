package ffmt // import "gopkg.in/ffmt.v1"

import (
	"io"
)

const colSym = ": "

// Space rune
var Space byte = ' '

// P go stlye display types
var defP = NewOptional(5, StlyeP, CanDefaultString|CanFilterDuplicate|CanRowSpan)

// Fp The go stlye friendly display types and data to writer
func Fp(w io.Writer, a ...interface{}) (int, error) {
	return defP.Fprint(w, a...)
}

// P The go stlye friendly display types and data
func P(a ...interface{}) (int, error) {
	return defP.Print(a...)
}

// Sp The go stlye friendly display types and data to string
func Sp(a ...interface{}) string {
	return defP.Sprint(a...)
}

// Puts go stlye
var defPuts = NewOptional(5, StlyePuts, CanDefaultString|CanFilterDuplicate|CanRowSpan)

// Fputs The go stlye friendly to writer
func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return defPuts.Fprint(w, a...)
}

// Puts The go stlye friendly display
func Puts(a ...interface{}) (int, error) {
	return defPuts.Print(a...)
}

// Sputs The go stlye friendly to string
func Sputs(a ...interface{}) string {
	return defPuts.Sprint(a...)
}

// Print go stlye
var defPrint = NewOptional(5, StlyePrint, CanDefaultString|CanFilterDuplicate|CanRowSpan)

// Fprint The go stlye friendly to writer
func Fprint(w io.Writer, a ...interface{}) (int, error) {
	return defPrint.Fprint(w, a...)
}

// Print The go stlye friendly display
func Print(a ...interface{}) (int, error) {
	return defPrint.Print(a...)
}

// Sprint The go stlye friendly to string
func Sprint(a ...interface{}) string {
	return defPrint.Sprint(a...)
}

// Pjson json stlye
var defPjson = NewOptional(20, StlyePjson, CanDefaultString|CanRowSpan)

// Fpjson The json stlye friendly display to writer
func Fpjson(w io.Writer, a ...interface{}) (int, error) {
	return defPjson.Fprint(w, a...)
}

// Pjson The json stlye friendly display
func Pjson(a ...interface{}) (int, error) {
	return defPjson.Print(a...)
}

// Spjson The json stlye friendly display to string
func Spjson(a ...interface{}) string {
	return defPjson.Sprint(a...)
}
