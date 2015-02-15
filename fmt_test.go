package ffmt

import (
	"testing"
)

func Test_fmt(t *testing.T) {
	Puts(Test1)
}

var Test1 = struct {
	id    int
	array []int
	maps  map[string]int
}{
	1,
	[]int{3, 4, 5, 6, 7},
	map[string]int{
		"aa": 21,
		"bb": 666,
	},
}
