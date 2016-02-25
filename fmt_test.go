package ffmt

import (
	"math"
	"time"

	"testing"
)

func TestFmtMap(t *testing.T) {
	bl := math.Pi / 10
	mm := map[interface{}]interface{}{}
	mm["i"] = "cos(i)"
	for i := float64(0); i <= 2*math.Pi; i += bl {
		mm[i] = math.Cos(i)
	}
	Print(mm)
}

func TestFmtSlice(t *testing.T) {
	bl := math.Pi / 10
	mc := []interface{}{}
	mc = append(mc, "\\", "i", "sin(i)")
	for i := float64(0); i <= 2*math.Pi; i += bl {
		mc = append(mc, i/bl+1, i, math.Sin(i))
	}
	Print(mc)
}

func TestFmtElse(t *testing.T) {
	Print(Test1)
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
}
