package main

type BinarySearchNode struct {
	key   int
	left  *BinarySearchNode
	right *BinarySearchNode
}

type BinarySearch struct {
	root *BinarySearchNode
}

func (b *BinarySearch) Insert(k int) {
	if b.root == nil {
		b.root = &BinarySearchNode{key: k}
		return
	}

	root := b.root
	for {
		if k == root.key {
			break
		} else if k > root.key {
			if root.right == nil {
				root.right = &BinarySearchNode{key: k}
				break
			}
			root = root.right
		} else if k < root.key {
			if root.left == nil {
				root.left = &BinarySearchNode{key: k}
				break
			}
			root = root.left
		}
	}
}

func (b *BinarySearch) Search(k int) bool {
	if b.root == nil {
		return false
	}

	root := b.root
	for {
		if k == root.key {
			return true
		} else if k > root.key {
			if root.right == nil {
				break
			}
			root = root.right
		} else if k < root.key {
			if root.left == nil {
				break
			}
			root = root.left
		}
	}

	return false
}
