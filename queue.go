package main

type Queue struct {
	elements []int
}

func (q *Queue) Enqueue(e int) {
	q.elements = append(q.elements, e)
}

func (q *Queue) Dequeue() (dequeueElement int) {
	dequeueElement = q.elements[0]
	q.elements = q.elements[1:]
	return
}
