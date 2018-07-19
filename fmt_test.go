package ffmt

import (
	"time"

	"testing"
)

func TestNewOptional(t *testing.T) {
	o := NewOptional(0, 0, 0)
	o.Print(testdata)
}

func TestPrint(t *testing.T) {
	Print(testdata)
}

func TestPuts(t *testing.T) {
	Puts(testdata)
}

func TestP(t *testing.T) {
	P(testdata)
}

func TestPjson(t *testing.T) {
	Pjson(testdata)
}

func TestD(t *testing.T) {
	D(testdata)
}

type bbb struct {
	A int
}

type T struct {
	bbb
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
	Fun    interface{}
}

var testdata = &T{
	bbb{},
	"Display a friendly fmt for golang",
	"你好",
	"hello all hello all hello all hello all hello all hello all ",
	"Display ",
	[]string{"hello", "world", "bey", "bey", "宽字符制表显示正常仅限等宽字体", "效率又降低了", "哈哈哈哈哈啊", "咳咳", "然而并没有什么卵用"},
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
		"一二三四五12345adcde": "1122334455",
		"鱼鱼鱼":             "yuyuyu",
	},
	true,
	time.Now(),
	nil,
	make(chan int, 10),
	Printf,
}
