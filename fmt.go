package ffmt

import "regexp"

var (
	WidthMax = 79
	LineMax  = 39
	reg      = regexp.MustCompile(`([\"\'][\w\s\.]+[^\"\'][\"\'])|([\w\.]+)|.`)
	//	regpro      = regexp.MustCompile(`[\[\{\(]`)
	//	regsuf      = regexp.MustCompile(`[\]\}\)]`)
	regpro   = regexp.MustCompile(`[\[\{]`)
	regsuf   = regexp.MustCompile(`[\]\}]`)
	regstrip = regexp.MustCompile(`[\s,]+`)
	//	regstrippro = regexp.MustCompile(`([\[\{\(])\s+`)
	//	regstripsuf = regexp.MustCompile(`\s+([\]\}\)])`)
	regstrippro = regexp.MustCompile(`([\[\{])\s+`)
	regstripsuf = regexp.MustCompile(`\s+([\]\}])`)
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
	b += "\n"
	return
}

func Fmt(a string) string {
	depth := -1
	width := 0
	line := 0
	ret := reg.ReplaceAllStringFunc(Strip(a), func(b string) (out string) {
		if regstrip.MatchString(b) {
			if width >= WidthMax || line >= LineMax {
				out = spacing(depth) + b
				width = depth
				line = 0
			} else {
				out = b
			}
		} else if regpro.MatchString(b) {
			depth++

			out = b + spacing(depth+1)
			width = depth
			line = 0
		} else if regsuf.MatchString(b) {
			out = spacing(depth) + b
			depth--
			width = WidthMax
		} else {
			out = b
		}
		line += len(b)
		width += len(out)
		return
	})
	return Trim(ret)
}
