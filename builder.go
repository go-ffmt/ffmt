package ffmt

import (
	"bytes"
	"fmt"
	"sync"
)

type builder interface {
	fmt.Stringer
	Write([]byte) (int, error)       // io.Writer
	WriteString(string) (int, error) // io.writeString
	WriteByte(byte) error            // io.ByteWriter
	WriteRune(rune) (int, error)
	Reset()
	Grow(int)
	Len() int
}

var poolBuilder = sync.Pool{
	New: newBuilder,
}

func getBuilder() builder {
	buf := poolBuilder.Get().(builder)
	return buf
}

func putBuilder(buf builder) {
	buf.Reset()
	poolBuilder.Put(buf)
}

func newBuilder() interface{} {
	return bytes.NewBuffer(nil)
}
