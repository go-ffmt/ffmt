package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	ffmt "gopkg.in/ffmt.v1"
)

var (
	w      = flag.Bool("w", false, "Write the changes to the file")
	prefix = flag.String("p", "//", "Prefix")
	split  = flag.String("s", ",", "Split rune")
)

func init() {
	flag.Usage = func() {
		w := os.Stdout
		fmt.Fprintf(w, "tablefmt:\n")
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
		err := format(file, *prefix, *split, *w)
		if err != nil {
			fmt.Println(err)
			flag.Usage()
			return
		}
	}
}

func format(file string, prefix, split string, w bool) error {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	ret := ffmt.TableText(string(b), prefix, split)
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
