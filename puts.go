package ffmt

import (
	"fmt"
	"io"
)

func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w,Sputs(a...))
}

func Puts(a ...interface{}) (int, error) {
	return fmt.Print(Sputs(a...))
}

func Sputs(a ...interface{}) string {
	return Fmt(ToString(a...))
}

