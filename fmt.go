package ffmt

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

var (
	WidthMax    = 0
	reg         = regexp.MustCompile(`([\"\'][\w\s\.]+[^\"\'][\"\'])|([\w\.]+)|.`)
	regpro      = regexp.MustCompile(`[\[\{\(]`)
	regsuf      = regexp.MustCompile(`[\]\}\)]`)
	regstrip    = regexp.MustCompile(`[\s,]+`)
	regstrippro = regexp.MustCompile(`([\[\{\(])\s+`)
	regstripsuf = regexp.MustCompile(`\s+([\]\}\)])`)
	regspar     = regexp.MustCompile(`\s+`)
	regtrim     = regexp.MustCompile(`\n\s+\n`)
	regcolon    = regexp.MustCompile(`:\s*`)
	regbracket  = regexp.MustCompile(`\n\s+([\[\{\(])`)
)

func spacing(depth int) string {
	b := []byte{'\n'}
	if depth > 0 {
		for i := 0; i != depth; i++ {
			b = append(b, ' ')
		}
	}
	return string(b)
}

func Strip(a string) (b string) {
	b = regspar.ReplaceAllString(a, " ")
	b = regstrippro.ReplaceAllString(b, "$1")
	b = regstripsuf.ReplaceAllString(b, "$1")
	return
}

func Trim(a string) (b string) {
	b = regtrim.ReplaceAllString(a, "\n")
	b = regcolon.ReplaceAllString(b, ":")
	//b = regbracket.ReplaceAllString(b, " $1")
	return
}

func Fmt(a string) string {
	depth := -1
	width := 0
	ret := reg.ReplaceAllStringFunc(Strip(a), func(b string) (out string) {
		if regstrip.MatchString(b) {
			if width >= WidthMax {
				out = spacing(depth) + b
			} else {
				out = b
			}
		} else if regpro.MatchString(b) {
			depth++
			out = b + spacing(depth+1)
			width = depth
		} else if regsuf.MatchString(b) {
			out = spacing(depth) + b
			depth--
			width = depth
		} else {
			out = b
		}
		width += len(out)
		return
	})
	return Trim(ret)
}

func Sputs(a ...interface{}) string {
	return Fmt(fmt.Sprintln(a...))
}

func Puts(a ...interface{}) (int, error) {
	return Fputs(os.Stdout, a...)
}

func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprintln(w, Sputs(a...))
}
