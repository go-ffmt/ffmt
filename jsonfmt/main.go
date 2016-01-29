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
	ret := ""
	if *file != "" {
		b, _ := ioutil.ReadFile(*file)
		var i interface{}
		json.Unmarshal(b, &i)
		ret = ffmt.Sjson(i)
		if *out != "" {
			ioutil.WriteFile(*out, []byte(ret), 0777)
		} else {
			fmt.Print(ret)
		}
	} else {
		flag.PrintDefaults()
	}
}
