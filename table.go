package ffmt

import (
	"bytes"
)

// 制表
func FmtTable(b [][]string) (ss []string) {
	maxs := []int{}
	for _, v1 := range b {
		for k, v2 := range v1 {
			if len(maxs) == k {
				maxs = append(maxs, 0)
			}
			if b := Biglen(v2); maxs[k] < b {
				maxs[k] = b
			}
		}
	}
	buf := bytes.NewBuffer(nil)
	for _, v1 := range b {
		buf.Reset()
		for k, v2 := range v1 {
			buf.WriteString(v2)
			ps := maxs[k] - Biglen(v2) + 1
			for i := 0; i != ps; i++ {
				buf.WriteByte(' ')
			}
		}
		ss = append(ss, buf.String())
	}
	return
}
