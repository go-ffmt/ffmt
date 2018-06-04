package main

import (
	"flag"
	"fmt"
	"io/ioutil"

	ffmt "gopkg.in/ffmt.v1"
)

var (
	file   = flag.String("f", "", "json file")
	out    = flag.String("o", "", "out file")
	prefix = flag.String("p", "//", "prefix")
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

	ret := ffmt.TableText(string(b), *prefix, *split)

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
