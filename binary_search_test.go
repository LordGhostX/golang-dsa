package main

import "testing"

func TestBinarySearchMethods(t *testing.T) {
	b := BinarySearch{}

	elementsToAdd := []int{4, 6, 10, 8, 7, 14}
	for _, v := range elementsToAdd {
		b.Insert(v)
	}

	testInput := []int{4, 3, 6, 10, 13, 7, 20}
	expected := []bool{true, false, true, true, false, true, false}
	for i := 0; i < len(testInput); i++ {
		if b.Search(testInput[i]) != expected[i] {
			t.Fail()
		}
	}
}
