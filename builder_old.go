// +build !go1.10

package ffmt

import (
	"bytes"
)

func newBuilder() interface{} {
	return bytes.NewBuffer(nil)
}
