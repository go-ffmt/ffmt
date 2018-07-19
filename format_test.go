package ffmt

import (
	"testing"
)

func TestFormat(t *testing.T) {
	if Format("hello {name}!", map[string]interface{}{
		"name": "world",
	}) != "hello world!" {
		t.Fail()
	}
}

func TestFormat2(t *testing.T) {
	if Format("hello {Name}!", struct{ Name string }{"world"}) != "hello world!" {
		t.Fail()
	}
}

func TestFormat3(t *testing.T) {
	if Format("hello {0}!", []string{"world"}) != "hello world!" {
		t.Fail()
	}
}
