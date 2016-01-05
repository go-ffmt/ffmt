package ffmt

import (
	"fmt"
	"io"
	"log"
)

func Fputs(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sputs(a...))
}

func Puts(a ...interface{}) (int, error) {
	return fmt.Print(Sputs(a...))
}

func Sputs(a ...interface{}) string {
	return Fmt(ToString(DepthMax, false, a...))
}

func Fp(w io.Writer, a ...interface{}) (int, error) {
	return fmt.Fprint(w, Sp(a...))
}

func P(a ...interface{}) (int, error) {
	return fmt.Print(Sp(a...))
}

func Sp(a ...interface{}) string {
	return Fmt(ToString(DepthMax, true, a...))
}

func Logs(a ...interface{}) {
	log.Print(Sputs(a...))
}

func Debug(a ...interface{}) {
	log.Print(Sp(a...))
}
