// +build go1.10

package ffmt

import (
	"strings"
)

func newBuilder() interface{} {
	return &strings.Builder{}
}
