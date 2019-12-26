package ffmt

import (
	"bytes"
	"fmt"
	"io"
	"sync"
)

type builder interface {
	fmt.Stringer
	io.Writer
	io.StringWriter
	io.ByteWriter
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
