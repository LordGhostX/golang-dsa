package main

import (
	"testing"
)

func TestGraph_Lookup(t *testing.T) {
	g := Graph{}

	elementsToAdd := [][2]string{{"school", "library"}, {"library", "church"}, {"mosque", "eatery"}, {"store", "park"}}
	for _, v := range elementsToAdd {
		g.Add(v[0], v[1])
	}

	testInput := [][2]string{{"school", "library"}, {"store", "park"}, {"park", "church"}}
	expected := []bool{true, true, false}
	for i := 0; i < len(testInput); i++ {
		if g.Lookup(testInput[i][0], testInput[i][1]) != expected[i] {
			t.Fail()
		}
	}
}

func TestGraph_Remove(t *testing.T) {
	g := Graph{}

	elementsToAdd := [][2]string{{"school", "library"}, {"field", "school"}, {"library", "church"}, {"mosque", "eatery"}, {"store", "park"}}
	for _, v := range elementsToAdd {
		g.Add(v[0], v[1])
	}

	elementsToRemove := []string{"school", "mosque", "store"}
	for _, v := range elementsToRemove {
		g.Remove(v)
	}

	testInput := [][2]string{{"school", "library"}, {"store", "park"}, {"library", "church"}}
	expected := []bool{false, false, true}
	for i := 0; i < len(testInput); i++ {
		if g.Lookup(testInput[i][0], testInput[i][1]) != expected[i] {
			t.Fail()
		}
	}
}
