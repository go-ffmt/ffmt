package ffmt

import (
	"time"

	"testing"
)

func Test_fmt(t *testing.T) {
	Json(Test1)
	Print(Test1)
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
	T      time.Time
	Inter  interface{}
	Chan   interface{}
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
	time.Now(),
	func(string) int { return 0 },
	make(chan int, 10),
}
