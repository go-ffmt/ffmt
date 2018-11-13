package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/ffmt.v1"
)

var (
	w = flag.Bool("w", false, "Write the changes to the file")
)

func init() {
	flag.Usage = func() {
		w := os.Stdout
		fmt.Fprintf(w, "jsonfmt:\n")
		fmt.Fprintf(w, "Usage:\n")
		fmt.Fprintf(w, "    %s [Options] file1 [filen ...]\n", os.Args[0])
		fmt.Fprintf(w, "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		return
	}
	for _, file := range args {
		err := format(file, *w)
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}
	}
}

func format(file string, w bool) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	var i interface{}
	err = json.Unmarshal(b, &i)
	if err != nil {
		return err
	}
	ret := ffmt.Spjson(i)
	if !w {
		fmt.Print(ret)
		return nil
	}

	err = ioutil.WriteFile(file, []byte(ret), 0666)
	if err != nil {
		return err
	}
	return nil
}
