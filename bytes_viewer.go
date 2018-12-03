package ffmt

import (
	"bytes"
	"fmt"
	"strconv"
	"unicode"
	"unsafe"
)

// BytesViewer bytes viewer
type BytesViewer []byte

// String returns view in hexadecimal
func (b BytesViewer) String() string {
	if len(b) == 0 {
		return invalid
	}
	const head = `
| Address  | Hex                                             | Text             |
| -------: | :---------------------------------------------- | :--------------- |
`
	const row = 16
	result := make([]byte, 0, len(head)/2*(len(b)/16+3))
	result = append(result, head...)
	for i := 0; i < len(b); i += row {
		result = append(result, "| "...)
		result = append(result, fmt.Sprintf("%08x", i)...)
		result = append(result, " | "...)

		k := i + row
		more := 0
		if k >= len(b) {
			more = k - len(b)
			k = len(b)
		}
		for j := i; j != k; j++ {
			if b[j] < 16 {
				result = append(result, '0')
			}
			result = strconv.AppendUint(result, uint64(b[j]), 16)
			result = append(result, ' ')
		}
		for j := 0; j != more; j++ {
			result = append(result, "   "...)
		}
		result = append(result, "| "...)
		buf := bytes.Map(func(r rune) rune {
			if unicode.IsSpace(r) {
				return ' '
			}
			return r
		}, b[i:k])
		result = append(result, buf...)
		for j := 0; j != more; j++ {
			result = append(result, ' ')
		}
		result = append(result, " |\n"...)
	}
	return *(*string)(unsafe.Pointer(&result))
}
