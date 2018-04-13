package main

import (
	"fmt"

	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	example()
}

func example() {
	m := struct {
		String string
		Int    int
		Slice  []int
		Map    map[string]interface{}
	}{
		"hello world",
		100,
		[]int{1, 2, 3, 4, 5, 6},
		map[string]interface{}{
			"A":  123,
			"BB": 456,
		},
	}

	fmt.Println(m) // fmt 默认输出
	/*
		{hello world 100 [1 2 3 4 5 6] map[BB:456 A:123]}
	*/

	ffmt.Puts(m) // 较为友好的输出
	/*
		{
		 String: "hello world"
		 Int:    100
		 Slice:  [
		  1 2 3
		  4 5 6
		 ]
		 Map: {
		  "A":  123
		  "BB": 456
		 }
		}
	*/

	ffmt.Print(m) // 同 Puts 但是字符串不加引号
	/*
		{
		 String: hello world
		 Int:    100
		 Slice:  [
		  1 2 3
		  4 5 6
		 ]
		 Map: {
		  A:  123
		  BB: 456
		 }
		}
	*/

	ffmt.P(m) // 友好格式化加上类型
	/*
		struct{
		 String: string(hello world)
		 Int:    int(100)
		 Slice:  slice[
		  int(1) int(2) int(3)
		  int(4) int(5) int(6)
		 ]
		 Map: map{
		  string(A):  int(123)
		  string(BB): int(456)
		 }
		}
	*/

	ffmt.Pjson(m) // 以 json 风格输出
	/*
		{
		 "Int": 100
		,"Map": {
		  "A":  123
		 ,"BB": 456
		 }
		,"Slice": [
		  1,2,3
		 ,4,5,6
		 ]
		,"String": "hello world"
		}
	*/

	m0 := ffmt.ToTable(m, m) // 按字段拆成表
	ffmt.Puts(m0)
	/*
		[
		 [
		  "String" "Int"
		  "Slice"  "Map"
		 ]
		 [
		  "hello world"   "100"
		  "[1 2 3 4 5 6]" "map[A:123 BB:456]"
		 ]
		]
	*/

	m1 := ffmt.FmtTable(m0) // [][]string 表格式化
	ffmt.Puts(m1)
	/*
		[
		 "String      Int Slice         Map               "
		 "hello world 100 [1 2 3 4 5 6] map[A:123 BB:456] "
		]
	*/

	ffmt.Mark("hello") // 标记输出位置
	/*
	   main.go:122  hello
	*/
}
