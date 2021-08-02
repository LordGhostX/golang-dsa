package main

type LinkedListNode struct {
	value    int
	nextLink *LinkedListNode
}

type LinkedList struct {
	head   *LinkedListNode
	length int
}

func (l *LinkedList) Append(v int) {
	if l.head == nil {
		l.head = &LinkedListNode{value: v}
	} else {
		currentLink := l.head
		for currentLink != nil {
			if currentLink.nextLink == nil {
				currentLink.nextLink = &LinkedListNode{value: v}
				break
			}
			currentLink = currentLink.nextLink
		}
	}
	l.length++
}

func (l *LinkedList) Prepend(v int) {
	l.head = &LinkedListNode{value: v, nextLink: l.head}
	l.length++
}

func (l *LinkedList) DeleteByValue(v int) {
	if l.head.value == v {
		l.head = l.head.nextLink
		l.length--
	} else {
		currentLink := l.head
		for currentLink.nextLink != nil {
			if currentLink.nextLink.value == v {
				currentLink.nextLink = currentLink.nextLink.nextLink
				l.length--
				break
			}
			currentLink = currentLink.nextLink
		}
	}
}
