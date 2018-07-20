package ffmt

import (
	"strings"
	"testing"
)

func TestAlign(t *testing.T) {
	b := `
{
 key: value
 slice: [
  hello
  world
  1
  2
 ]	
}
`

	out := `
{
 key:   value
 slice: [
  hello world
  1     2
 ]	
}
`
	if strings.TrimSpace(Align(b)) != strings.TrimSpace(out) {
		t.Fail()
	}
}
