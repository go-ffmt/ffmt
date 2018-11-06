package ffmt

import (
	"strings"
	"testing"
)

func TestBytesViewer(t *testing.T) {
	tests := []struct {
		b    BytesViewer
		want string
	}{
		{BytesViewer("Hello world!"), `| Address  | Hex                                             | Text             |
| -------: | :---------------------------------------------- | :--------------- |
| 00000000 | 48 65 6c 6c 6f 20 77 6f 72 6c 64 21             | Hello world!     |`},
	}
	for _, tt := range tests {
		if got := tt.b.String(); strings.TrimSpace(got) != strings.TrimSpace(tt.want) {
			t.Fail()
		}
	}
}
