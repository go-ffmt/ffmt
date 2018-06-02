package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"strings"

	"gopkg.in/ffmt.v1"
)

var (
	file = flag.String("f", "", "json file")
	out  = flag.String("o", "", "out file")
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
	var i interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		fmt.Println(err)
		flag.PrintDefaults()
		return
	}
	ret := ffmt.Spjson(i)
	if *out == "" {
		fmt.Print(ret)
		return
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
