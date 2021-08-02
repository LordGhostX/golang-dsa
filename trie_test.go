package main

import (
	"testing"
)

func TestTrieMethods(t *testing.T) {
	tr := Trie{}

	elementsToAdd := []string{"hello", "world", "fish", "dog", "house"}
	for _, v := range elementsToAdd {
		tr.Insert(v)
	}

	testInput := []string{"world", "cat", "house", "dog", "school", "fire", "fis"}
	expected := []bool{true, false, true, true, false, false, false}
	for i := 0; i < len(testInput); i++ {
		if tr.Search(testInput[i]) != expected[i] {
			t.Fail()
		}
	}
}
