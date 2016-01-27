package ffmt

import "regexp"

var (
	//	WidthMax = 0
	//	LineMax  = 0
	WidthMax    = 79
	LineMax     = 39
	reg         = regexp.MustCompile(`([\w\.]+)|('[^']*')|("[^"]*")|(<[^<>]*>)|(\([^\(]*\))|.`)
	regpro      = regexp.MustCompile(`[\[\{]`)
	regsuf      = regexp.MustCompile(`[\]\}]`)
	regstrip    = regexp.MustCompile(`[\s,]+`)                             // json字符串前缀逗好
	regstrippro = regexp.MustCompile(`([\(\[\{<])\s*`)                     // 设置左括号右边空一格
	regstripsuf = regexp.MustCompile(`\s*([\)\]\}>])`)                     // 设置右括号左边空一格
	regspar     = regexp.MustCompile(`\s+`)                                // 删除多余空格
	regtrim     = regexp.MustCompile(`[\n\s*]+\n`)                         // 删除多余行
	regcolon    = regexp.MustCompile(`:\s*`)                               // 冒号后面空一格
	regbracket  = regexp.MustCompile(`([^\[\{\(\<\s,\"])\s*([\[\{\(\<:])`) // 左括号顶到空一格
	regempty    = regexp.MustCompile(`([\[\{\(])\s+([\]\}\)])`)            // 如果 括号之间没有东西 去除括号之间的空格
)

func spac(depth int) string {
	b := []byte{}
	if depth > 0 {
		for i := 0; i != depth; i++ {
			b = append(b, ' ')
		}
	}
	return string(b)
}

func spacing(depth int) string {
	return "\n" + spac(depth)
}

func strip(a string) (b string) {
	b = a
	b = regspar.ReplaceAllString(b, " ")
	b = regstrippro.ReplaceAllString(b, "$1")
	b = regstripsuf.ReplaceAllString(b, "$1")
	return
}

func trim(a string) (b string) {
	b = a
	b = regcolon.ReplaceAllString(b, ": ")
	b = regempty.ReplaceAllString(b, "$1$2")
	b = regbracket.ReplaceAllString(b, "$1$2")
	return
}

func moveSpac(a string) (b string) {
	b = a
	b += "\n"
	b = regtrim.ReplaceAllString(b, "\n")
	return
}

func fmts(a string, widthMax, lineMax int) string {
	depth := -1
	width := 0
	line := 0
	ret := reg.ReplaceAllStringFunc(strip(a), func(b string) (out string) {
		if regstrip.MatchString(b) {
			if width >= widthMax || line >= lineMax {
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
			width = widthMax
		} else {
			out = b
		}
		line += len(b)
		width += len(out)
		return
	})
	return trim(ret)
}
