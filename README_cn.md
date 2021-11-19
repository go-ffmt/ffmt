# golang 美化数据展示

[![Go Report Card](https://goreportcard.com/badge/github.com/go-ffmt/ffmt)](https://goreportcard.com/report/github.com/go-ffmt/ffmt)
[![GoDoc](https://pkg.go.dev/badge/github.com/gopkg.in/ffmt.v1)](https://pkg.go.dev/gopkg.in/ffmt.v1)
[![GitHub license](https://img.shields.io/github/license/go-ffmt/ffmt.svg)](https://github.com/go-ffmt/ffmt/blob/master/LICENSE)
[![gocover.io](https://gocover.io/_badge/github.com/go-ffmt/ffmt)](https://gocover.io/github.com/go-ffmt/ffmt)

- [English](https://github.com/go-ffmt/ffmt/blob/master/README.md)
- [简体中文](https://github.com/go-ffmt/ffmt/blob/master/README_cn.md)

## 用法

[示例](https://github.com/go-ffmt/ffmt/blob/master/examples/main.go)

``` golang
package main

import (
	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	example()
}

type mt struct {
	String string
	Int    int
	Slice  []int
	Map    map[string]interface{}
}

func example() {
	m := mt{
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
		main.mt{
		 String: string("hello world")
		 Int:    int(100)
		 Slice:  []int[
		  int(1) int(2) int(3)
		  int(4) int(5) int(6)
		 ]
		 Map: map[string]interface {}{
		  string("A"):  int(123)
		  string("BB"): int(456)
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
		main.go:124  hello
	*/

	ffmt.Print(ffmt.BytesViewer("Hello world! Hello All!"))
	/*
  |  Address | Hex                                             | Text             |
  | -------: | :---------------------------------------------- | :--------------- |
  | 00000000 | 48 65 6c 6c 6f 20 77 6f 72 6c 64 21 20 48 65 6c | Hello world! Hel |
  | 00000010 | 6c 6f 20 41 6c 6c 21                            | lo All!          |
	*/
}
```

## 赞助

[![jetbrains](https://www.jetbrains.com/shop/static/images/jetbrains-logo-inv.svg)](https://www.jetbrains.com/shop/eform/opensource)

## 许可证

软包根据MIT License。有关完整的许可证文本，请参阅[LICENSE](https://github.com/go-ffmt/ffmt/blob/master/LICENSE)。
