package main

import "testing"

var Debug bool = false

func TestIsOne(t *testing.T) {
	i := 1
	if Debug {
		t.Skip("スキップする")
	}
	v := IsOne(i)
	if !v {
		t.Errorf("%v != %v", 1, i)
	}
}
