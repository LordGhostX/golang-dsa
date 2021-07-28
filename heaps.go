package main

type Heap struct {
	elements []int
}

func (h *Heap) Insert(e int) {
	h.elements = append(h.elements, e)
	h.heapUpwards()
}

func (h *Heap) heapUpwards() {
	lastIndex := len(h.elements) - 1
	parent := findParent(lastIndex)
	for h.elements[lastIndex] > h.elements[parent] {
		h.elements[lastIndex], h.elements[parent] = h.elements[parent], h.elements[lastIndex]
		lastIndex = parent
		parent = findParent(lastIndex)
	}
}

func (h *Heap) Remove() (e int) {
	e = h.elements[0]
	lastIndex := len(h.elements) - 1

	h.elements[0] = h.elements[lastIndex]
	h.elements = h.elements[:lastIndex]
	h.heapDownwards()
	return
}

func (h *Heap) heapDownwards() {
	lastIndex, depth := 0, len(h.elements)-1
	leftChild, rightChild := findChildren(lastIndex, depth)
	for leftChild != -1 || rightChild != -1 {
		biggestChild := h.findBiggestChild([]int{lastIndex, leftChild, rightChild})
		if biggestChild == lastIndex {
			break
		}
		h.elements[lastIndex], h.elements[biggestChild] = h.elements[biggestChild], h.elements[lastIndex]
		lastIndex = biggestChild
		leftChild, rightChild = findChildren(lastIndex, depth)
	}
}

func (h *Heap) findBiggestChild(children []int) int {
	biggestChild := children[0]
	for _, v := range children {
		if v == -1 {
			continue
		}
		if h.elements[v] > h.elements[biggestChild] {
			biggestChild = v
		}
	}
	return biggestChild
}

func findParent(i int) int {
	var parent int
	if i%2 == 0 {
		parent = (i - 2) / 2
	} else {
		parent = (i - 1) / 2
	}

	if parent < 0 {
		return 0
	}
	return parent
}

func findChildren(i, depth int) (int, int) {
	leftChild, rightChild := i*2+1, i*2+2

	if leftChild > depth {
		leftChild = -1
	}
	if rightChild > depth {
		rightChild = -1
	}
	return leftChild, rightChild
}
