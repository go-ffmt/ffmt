package ffmt

import (
	"time"

	"testing"
)

func Test_fmt(t *testing.T) {
	Test1.TTT = &Test1
	Json(Test1)
	Print(Test1)
	Puts(Test1)
	P(Test1)
}

var Test1 = struct {
	Msg  string
	Msg2 string
	Msg3 string
	msg  string
	Msgs []string
	Stru []struct {
		Msg string
		AA  [8]int
	}

	Floats [6]float32
	Ints   [][]int
	Maps   map[string]string
	B      bool
	T      time.Time
	TTT    interface{}
	Chan   interface{}
}{

	"Display a friendly fmt for golang",
	"你好",
	"hello all hello all hello all hello all hello all hello all ",
	"Display ",
	[]string{"hello", "world", "bey", "bey"},
	[]struct {
		Msg string
		AA  [8]int
	}{{}, {
		"Test",
		[8]int{2222, 3333},
	}},
	[6]float32{2.1, 3.3},
	[][]int{{1, 4, 5, 1, 4, 5, 6, 11999, 0}, {3}, {}},
	map[string]string{
		"aa": "hi world",
		"bb": "bye world",
	},
	true,
	time.Now(),
	nil,
	make(chan int, 10),
}
