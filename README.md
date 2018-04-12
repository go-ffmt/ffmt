# Friendly data formatting for golang

[![Build Status](https://travis-ci.org/go-ffmt/ffmt.svg?branch=master)](https://travis-ci.org/go-ffmt/ffmt)
[![Go Report Card](https://goreportcard.com/badge/gopkg.in/ffmt.v1)](https://goreportcard.com/report/gopkg.in/ffmt.v1)
[![GoDoc](https://godoc.org/gopkg.in/ffmt.v1?status.svg)](https://godoc.org/gopkg.in/ffmt.v1)
[![GitHub license](https://img.shields.io/github/license/go-ffmt/ffmt.svg)](https://github.com/go-ffmt/ffmt/blob/master/LICENSE)

 - [English](./README.md)
 - [简体中文](./README_cn.md)

## Install

``` shell
go get -u -v gopkg.in/ffmt.v1
```

## Usage

[API Documentation](http://godoc.org/gopkg.in/ffmt.v1)

[Examples](./examples/main.go)

``` golang
package main

import (
	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	example()
}

func example() {
	m := map[string]interface{}{
		"hello": "w",
		"A": []int{
			1, 2, 3, 4, 5, 6,
		},
	}

	ffmt.Puts(m)
	/*
	   {
	    "A": [
	     1 2 3
	     4 5 6
	    ]
	    "hello": "w"
	   }
	*/

	ffmt.P(m)
	/*
	   map{
	    string(A): slice[
	     int(1) int(2) int(3)
	     int(4) int(5) int(6)
	    ]
	    string(hello): string(w)
	   }
	*/

	ffmt.Pjson(m)
	/*
	   {
	    "A": [
	     1,2,3
	    ,4,5,6
	    ]
	   ,"hello": "w"
	   }
	*/

	m0 := ffmt.ToTable(m, m)
	ffmt.Puts(m0)
	/*
	   [
	    [
	     "A"
	     "hello"
	    ]
	    [
	     "[1 2 3 4 5 6]"
	     "w"
	    ]
	   ]
	*/

	m1 := ffmt.FmtTable(m0)
	ffmt.Puts(m1)
	/*
	   [
	    "A             hello "
	    "[1 2 3 4 5 6] w     "
	   ]
	*/

	ffmt.Mark("hello")
	/*
	   main.go:76  hello
	*/
}

```




## License

Pouch is licensed under the MIT License. See [LICENSE](./LICENSE) for the full license text.
