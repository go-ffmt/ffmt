package ffmt

import (
	"bytes"
	"sync"
)

var pool = sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(nil)
	},
}

func getBuilder() *bytes.Buffer {
	buf := pool.Get().(*bytes.Buffer)
	buf.Reset()
	return buf
}

func putBuilder(buf *bytes.Buffer) {
	pool.Put(buf)
}
