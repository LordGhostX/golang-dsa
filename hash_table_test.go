package main

import (
	"testing"
)

func TestHashTableSearch(t *testing.T) {
	h := HashTable{size: 5}

	elementsToAdd := []string{"hello", "house", "fire", "school"}
	for _, v := range elementsToAdd {
		h.Insert(v)
	}

	testInput := []string{"hello", "world", "school", "vehicle", "danger"}
	expected := []bool{true, false, true, false, false}
	for i := 0; i < len(testInput); i++ {
		if h.Search(testInput[i]) != expected[i] {
			t.Fail()
		}
	}
}

func TestHashTableDelete(t *testing.T) {
	h := HashTable{size: 5}

	elementsToAdd := []string{"hello", "world", "house", "fire", "school"}
	for _, v := range elementsToAdd {
		h.Insert(v)
	}

	elementsToDelete := []string{"hello", "world", "fire", "fish"}
	for i := 0; i < len(elementsToDelete); i++ {
		h.Delete(elementsToDelete[i])
	}

	testInput := []string{"hello", "world", "house", "fire", "school"}
	expected := []bool{false, false, true, false, true}
	for i := 0; i < len(testInput); i++ {
		if h.Search(testInput[i]) != expected[i] {
			t.Fail()
		}
	}
}