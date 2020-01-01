package ffmt

import (
	"bytes"
	"sync"
)

type builder interface {
	String() string
	Write([]byte) (int, error)       // io.Writer
	WriteString(string) (int, error) // io.WriteString
	WriteByte(byte) error            // io.ByteWriter
	Reset()
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
	const malloc = 1024
	return bytes.NewBuffer(make([]byte, 0, malloc))
}
