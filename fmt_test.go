package ffmt

import (
	"encoding/json"
	"testing"
)

func Test_fmt(t *testing.T) {
	js, _ := json.Marshal(Test1)
	Puts(Sputs(string(js)))
	Puts(Sputs(Test1))
}

var Test1 = struct {
	Msg  string
	Stru struct {
		Msg string
		AA  []int
	}
	Floats []float32
	Ints   [][]int
	Maps   map[string]string
}{
	"Display a friendly fmt for golang",
	struct {
		Msg string
		AA  []int
	}{
		"Test",
		[]int{2222, 3333},
	},
	[]float32{2.1, 3.3},
	[][]int{{1, 4}, {3}},
	map[string]string{
		"aa": "hi world",
		"bb": "bye world",
	},
}
