package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"

	ffmt "gopkg.in/ffmt.v1"
)

var (
	file   = flag.String("f", "", "json file")
	out    = flag.String("o", "", "out file")
	prefix = flag.String("p", "", "prefix")
	split  = flag.String("s", ",", "Split rune")
)

func init() {
	flag.Parse()
}

func main() {
	if *file == "" {
		flag.PrintDefaults()
		return
	}
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
		return
	}
	rows := []string{}
	table := [][]string{}
	for _, v := range strings.Split(string(b), "\n") {
		if *prefix != "" && !strings.HasPrefix(v, *prefix) {
			if len(table) != 0 {
				for _, v := range ffmt.FmtTable(table) {
					ffmt.Mark(v)
					rows = append(rows, v)
				}
			}
			rows = append(rows, v)
			continue
		}

		row := []string{}
		ss := strings.Split(v, *split)
		for i, col := range ss {
			if i == 0 {
				row = append(row, strings.TrimRightFunc(col, unicode.IsSpace))
			} else {
				row = append(row, strings.TrimSpace(col))
			}
			if i != len(ss)-1 {
				row[i] = row[i] + *split
			}
		}
		table = append(table, row)
	}
	if len(table) != 0 {
		for _, v := range ffmt.FmtTable(table) {
			ffmt.Mark(v)
			rows = append(rows, v)
		}
	}

	ret := strings.Join(rows, "\n")
	if *out != "" {
		err = ioutil.WriteFile(*out, []byte(ret), 0666)
		if err != nil {
			fmt.Println(err)
			flag.PrintDefaults()
			return
		}
	} else {
		fmt.Print(ret)
	}
	return
}
