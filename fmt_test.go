package ffmt

import (
	"time"

	"encoding/json"
	"testing"
)

func Test_fmt(t *testing.T) {
	js, _ := json.Marshal(Test1)
	Print(string(js))
	Puts(Test1)
	P(Test1)
}

var Test1 = struct {
	Msg  string
	Msg2 string
	Msg3 string
	Msg4 string
	msg  string
	Stru []struct {
		Msg string
		AA  [5]int
	}

	Floats [5]float32
	Ints   [][]int
	Maps   map[string]string
	B      bool
}{

	"Display a friendly fmt for golang",
	"你好",
	"",
	"hello all hello all hello all hello all hello all hello all ",
	"Display ",
	[]struct {
		Msg string
		AA  [5]int
	}{{}, {
		"Test",
		[5]int{2222, 3333},
	}},
	[5]float32{2.1, 3.3},
	[][]int{{1, 4}, {3}},
	map[string]string{
		"aa": "hi world",
		"bb": "bye world",
	},
	true,
}

func Test_Now(t *testing.T) {
	P(time.Now())
	P()
}
