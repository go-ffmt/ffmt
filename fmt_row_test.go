package ffmt

import (
	"bytes"
	"strings"
	"testing"
)

var rowtestdata = [][]int{}

func init() {
	for i := 0; i != 20; i++ {
		rowtestdata = append(rowtestdata, make([]int, i))
	}
}
func TestRowPrint(t *testing.T) {
	buf := bytes.NewBuffer(nil)
	for _, v := range rowtestdata {
		Fprint(buf, v)
	}

	if strings.TrimSpace(buf.String()) != strings.TrimSpace(rowtestdataout) {
		t.Fail()
	}
}

var rowtestdataout = `
[
]
[
 0
]
[
 0
 0
]
[
 0
 0
 0
]
[
 0 0
 0 0
]
[
 0
 0
 0
 0
 0
]
[
 0 0 0
 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0
 0 0 0 0
]
[
 0 0 0
 0 0 0
 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0
 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0
 0 0 0 0 0 0 0
]
[
 0 0 0 0 0
 0 0 0 0 0
 0 0 0 0 0
]
[
 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
[
 0 0 0 0 0 0 0 0 0
 0 0 0 0 0 0 0 0 0
]
[
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
 0
]
`
