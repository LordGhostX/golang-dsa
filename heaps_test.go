package main

import (
	"testing"
)

func TestHeapMethods(t *testing.T) {
	h := Heap{}

	elementsToAdd := []int{10, 5, 2, 15, 8, 20, 3, 11, 50, 23, 25, 30}
	for _, v := range elementsToAdd {
		h.Insert(v)
	}

	expected := []int{50, 30, 25, 23, 20, 15, 11, 10, 8, 5, 3, 2}
	for _, v := range expected {
		if v != h.Remove() {
			t.Fail()
		}
	}
}
