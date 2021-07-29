package main

import "testing"

func TestStackMethods(t *testing.T) {
	s := Stack{}

	elementsToAdd := []int{2, 3, 5, 6, 7, 9, 3}
	for _, v := range elementsToAdd {
		s.Push(v)
	}

	expected := []int{3, 9, 7, 6, 5, 3, 2}
	for _, v := range expected {
		if v != s.Pop() {
			t.Fail()
		}
	}
}
