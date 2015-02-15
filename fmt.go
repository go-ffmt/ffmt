package ffmt

import (
	f "fmt"
	"io"
	"os"
	"regexp"
)

var (
	WidthMax = 0
	regpro   = regexp.MustCompile(`([\[\{\(])\s*`)
	regsuf   = regexp.MustCompile(`([\]\}\)])\s*`)
	reg      = regexp.MustCompile(`([\[\]\{\}\(\)[\"\']\s,])[,]\s*|.`)
	regspac  = regexp.MustCompile(`([\s,])\s*`)
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

func Fmt(a string) string {
	depth := 0
	width := 0
	o := reg.ReplaceAllStringFunc(a, func(b string) (out string) {
		if regpro.MatchString(b) {
			depth++
			width = 0
			out = f.Sprintf("%s%s", b, spacing(depth+1))
		} else if regsuf.MatchString(b) {
			out = f.Sprintf("%s%s", spacing(depth), b)
			depth--
			width = 0
		} else if regspac.MatchString(b) && width > WidthMax {
			out = f.Sprintf("%s%s", spacing(depth), b)
		} else {
			out = b
		}
		width += len(out)
		return
	})
	return o
}

func Sputs(a ...interface{}) string {
	return Fmt(f.Sprintln(a...))
}

func Puts(a ...interface{}) (n int, err error) {
	return Fputs(os.Stdout, a...)
}

func Fputs(w io.Writer, a ...interface{}) (n int, err error) {
	return f.Fprintln(w, Sputs(a...))
}
