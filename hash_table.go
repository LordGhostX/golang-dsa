package main

type HashTableBucket struct {
	value    string
	nextLink *HashTableBucket
}

type HashTable struct {
	size   int
	blocks []*HashTableBucket
}

func (h *HashTable) createBlocks() {
	if h.blocks == nil {
		h.blocks = make([]*HashTableBucket, h.size)
	}
}

func (h *HashTable) findIndex(k string) int {
	index := 0
	for i := range k {
		index += int(k[i])
	}
	return index % h.size
}

func (h *HashTable) Insert(v string) {
	h.createBlocks()
	index := h.findIndex(v)
	h.blocks[index] = &HashTableBucket{value: v, nextLink: h.blocks[index]}
}

func (h *HashTable) Search(v string) bool {
	h.createBlocks()
	index := h.findIndex(v)
	if h.blocks[index] == nil {
		return false
	} else {
		currentBucket := h.blocks[index]
		for currentBucket != nil {
			if currentBucket.value == v {
				return true
			}
			currentBucket = currentBucket.nextLink
		}
	}
	return false
}

func (h *HashTable) Delete(v string) {
	h.createBlocks()
	index := h.findIndex(v)
	if h.blocks[index] == nil {
		return
	} else if h.blocks[index].value == v {
		h.blocks[index] = h.blocks[index].nextLink
	} else {
		currentBucket := h.blocks[index]
		for currentBucket.nextLink != nil {
			if currentBucket.nextLink.value == v {
				currentBucket.nextLink = currentBucket.nextLink.nextLink
				break
			}
			currentBucket = currentBucket.nextLink
		}
	}
}
