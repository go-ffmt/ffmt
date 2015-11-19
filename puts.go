package ffmt

import (
	"fmt"
	"io"
	"os"
)

func Sputs(a ...interface{}) string {
	return Fmt(ToString(a...))
}

func Puts(a ...interface{}) (int, error) {
	return Fputs(os.Stdout, a...)
}

func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprintln(w, Sputs(a...))
}
