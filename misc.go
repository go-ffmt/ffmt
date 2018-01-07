package ffmt

import (
	"bytes"
	"fmt"
	"unicode"

	"sync"
)

// 默认
var Printf = fmt.Printf
var Println = fmt.Println

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
}

var BigWord = unicode.Scripts["Han"]

func Biglen(str string) int {
	i := 0
	for _, v := range str {
		if unicode.Is(BigWord, v) {
			i += 2
		} else {
			i++
		}
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
