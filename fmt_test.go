package ffmt

import (
	"time"

	"encoding/json"
	"testing"
)

func Test_fmt(t *testing.T) {
	js, _ := json.Marshal(Test1)
	Puts(Test1, string(js))

}

var Test1 = struct {
	Msg  string
	msg  string
	Stru []struct {
		Msg string
		AA  [20]int
	}

	Floats [20]float32
	Ints   [][]int
	Maps   map[string]string
	B      bool
}{
	"Display a friendly fmt for golang",
	"Display ",
	[]struct {
		Msg string
		AA  [20]int
	}{{}, {
		"Test",
		[20]int{2222, 3333},
	}},
	[20]float32{2.1, 3.3},
	[][]int{{1, 4}, {3}},
	map[string]string{
		"aa": "hi world",
		"bb": "bye world",
	},
	true,
}

func Test_Now(t *testing.T) {
	Puts(Test_Now)
	Puts(make(chan int))
	Puts(time.Now())
}
