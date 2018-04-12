package ffmt

import (
	"fmt"
	"unicode/utf8"
)

// Printf fmt.Printf
var Printf = fmt.Printf

// Println fmt.Println
var Println = fmt.Println

func runeWidth(r rune) int {
	switch {
	case r == utf8.RuneError || r < '\x20':
		return 0

	case '\x20' <= r && r < '\u2000':
		return 1

	case '\u2000' <= r && r < '\uFF61':
		return 2

	case '\uFF61' <= r && r < '\uFFA0':
		return 1

	case '\uFFA0' <= r:
		return 2
	}

	return 0
}

func strLen(str string) int {
	i := 0
	for _, v := range str {
		i += runeWidth(v)
	}
	return i
}

func spac(depth int) string {
	b := []byte{}
	if depth > 0 {
		for i := 0; i != depth; i++ {
			b = append(b, Space)
		}
	}
	return string(b)
}

func spacing(depth int) string {
	return "\n" + spac(depth-1)
}
