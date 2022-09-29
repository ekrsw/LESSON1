package foo

import "testing"

var Debug bool = false

func TestReturnMin(t *testing.T) {
	if Debug {
		t.Skip("スキップする")
	}
	m := ReturnMin()
	if m != 1 {
		t.Errorf("%v !+ %v", m, 1)
	}
}
