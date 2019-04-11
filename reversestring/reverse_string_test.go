package reversestring

import "testing"

var str = "Hello, world!"
var reversedStr = "!dlrow ,olleH"

func TestReverseBySlice(t *testing.T) {
	s := ReverseBySlice(str)
	if s != "!dlrow ,olleH" {
		t.Errorf("Expected reseverse string %s, but got %s.", str, s)
	}
}

func TestReverseByForloop(t *testing.T) {
	s := ReverseByForloop(str)

	if s != "!dlrow ,olleH" {
		t.Errorf("Expected reseverse string %s, but got %s.", str, s)
	}
}
