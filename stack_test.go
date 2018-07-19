package ffmt

import "testing"

func TestMarkStack(t *testing.T) {
	MarkStack(2, "hello")
}

func TestMark(t *testing.T) {
	Mark("hello")
}

func TestMarkStackFull(t *testing.T) {
	MarkStackFull()
}
