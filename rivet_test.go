package rivet

import "testing"

func TestConcat(t *testing.T) {
	c := Concat("a", "b", "c")
	if c != "abc" {
		t.Fail()
	}
}
