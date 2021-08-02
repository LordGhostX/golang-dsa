package main

import "testing"

func TestQueue_Methods(t *testing.T) {
	q := Queue{}

	elementsToAdd := []int{2, 3, 5, 6, 7, 9, 3}
	for _, v := range elementsToAdd {
		q.Enqueue(v)
	}

	expected := []int{2, 3, 5, 6, 7, 9, 3}
	for _, v := range expected {
		if v != q.Dequeue() {
			t.Fail()
		}
	}
}
