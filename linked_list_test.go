package main

import (
	"testing"
)

func TestLinkedList_Append(t *testing.T) {
	l := LinkedList{}

	for i := 0; i < 10; i++ {
		l.Append(i * i)
		if l.length != i+1 {
			t.Fail()
		}
	}
}

func TestLinkedList_Prepend(t *testing.T) {
	l := LinkedList{}

	for i := 0; i < 10; i++ {
		l.Prepend(i * i)
		if l.head.value != i*i {
			t.Fail()
		}
	}
}

func TestLinkedList_Delete(t *testing.T) {
	l := LinkedList{}

	elementsToAdd := []int{10, 5, 8, 3, 9, 20}
	for _, v := range elementsToAdd {
		l.Append(v)
	}

	testInput := []int{4, 3, 5, 10, 13, 9, 7}
	expected := []int{6, 5, 4, 3, 3, 2, 2}
	for i := 0; i < len(testInput); i++ {
		l.Delete(testInput[i])
		if l.length != expected[i] {
			t.Fail()
		}
	}
}
