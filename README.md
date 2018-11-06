# Golang beautify data display for Humans

[![Build Status](https://travis-ci.org/go-ffmt/ffmt.svg?branch=master)](https://travis-ci.org/go-ffmt/ffmt)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ffmt/ffmt)](https://goreportcard.com/report/github.com/go-ffmt/ffmt)
[![GoDoc](https://godoc.org/github.com/go-ffmt/ffmt?status.svg)](https://godoc.org/github.com/go-ffmt/ffmt)
[![GitHub license](https://img.shields.io/github/license/go-ffmt/ffmt.svg)](https://github.com/go-ffmt/ffmt/blob/master/LICENSE)
[![cover.run](https://cover.run/go/gopkg.in/ffmt.v1.svg?style=flat&tag=golang-1.10)](https://cover.run/go?tag=golang-1.10&repo=gopkg.in%2Fffmt.v1)

- [English](https://github.com/go-ffmt/ffmt/blob/master/README.md)
- [简体中文](https://github.com/go-ffmt/ffmt/blob/master/README_cn.md)

## Install

``` shell
# Stable version
go get -u -v gopkg.in/ffmt.v1

# Latest version
go get -u -v github.com/go-ffmt/ffmt
```

## Usage

[API Documentation](https://godoc.org/github.com/go-ffmt/ffmt)

[Examples](https://github.com/go-ffmt/ffmt/blob/master/examples/main.go)

``` golang
package main

import (
	ffmt "github.com/go-ffmt/ffmt"
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

	fmt.Println(m) // fmt the default formatting.
	/*
		{hello world 100 [1 2 3 4 5 6] map[BB:456 A:123]}
	*/

	ffmt.Puts(m) // More friendly formatting.
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

	ffmt.Print(m) // Same "Puts" but String unadded '"'.
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

	ffmt.P(m) // Format data and types.
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

	ffmt.Pjson(m) // Format it in json style.
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

	m0 := ffmt.ToTable(m, m) // Break the fields into tables.
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

	m1 := ffmt.FmtTable(m0) // [][]string Table format.
	ffmt.Puts(m1)
	/*
		[
		 "String      Int Slice         Map               "
		 "hello world 100 [1 2 3 4 5 6] map[A:123 BB:456] "
		]
	*/

	ffmt.Mark("hello") // Mark position.
	/*
		main.go:124  hello
	*/

	ffmt.Print(ffmt.BytesViewer("Hello world! Hello All!"))
	/*
		| Address  | Hex                                             | Text             |
		| -------: | :---------------------------------------------- | :--------------- |
		| 00000000 | 48 65 6c 6c 6f 20 77 6f 72 6c 64 21 20 48 65 6c | Hello world! Hel |
		| 00000010 | 6c 6f 20 41 6c 6c 21                            | lo All!          |
	*/
}


```

## License

Pouch is licensed under the MIT License. See [LICENSE](https://github.com/go-ffmt/ffmt/blob/master/LICENSE) for the full license text.
