package main

type Stack struct {
	elements []int
}

func (s *Stack) Push(e int) {
	s.elements = append(s.elements, e)
}

func (s *Stack) Pop() (popElement int) {
	stackLen := len(s.elements) - 1
	popElement = s.elements[stackLen]
	s.elements = s.elements[:stackLen]
	return
}
