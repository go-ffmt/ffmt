package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/wzshiming/ffmt"
)

func main() {
	file := flag.String("f", "", "json file")
	out := flag.String("o", "", "out file")
	flag.Parse()
	if *file == "" {
		flag.PrintDefaults()
		return
	}
	b, err := ioutil.ReadFile(*file)
	if err != nil {
		flag.PrintDefaults()
		return
	}
	var i interface{}
	json.Unmarshal(b, &i)
	ret := ffmt.Spjson(i)
	if *out == "" {
		fmt.Print(ret)
		return
	}
	err = ioutil.WriteFile(*out, []byte(ret), 0777)
	if err != nil {
		flag.PrintDefaults()
		return
	}
	return
}
