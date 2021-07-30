package main

import (
	"fmt"
	"strings"
)

type Node struct {
	children [26]*Node
	isEnd    bool
}

type Trie struct {
	root [26]*Node
}

func (t *Trie) Insert(s string) {
	sLen := len(s)
	currentBranch := &t.root
	for i := 0; i < sLen; i++ {
		charIndex := getCharIndex(s[i : i+1])
		if currentBranch[charIndex] == nil {
			currentBranch[charIndex] = &Node{}
		}
		if i == sLen-1 {
			currentBranch[charIndex].isEnd = true
		}
		currentBranch = &currentBranch[charIndex].children
	}
}

func (t *Trie) Search(s string) bool {
	sLen := len(s)
	currentBranch := &t.root
	for i := 0; i < sLen; i++ {
		charIndex := getCharIndex(s[i : i+1])
		if currentBranch[charIndex] == nil {
			return false
		}
		if i == sLen-1 {
			if currentBranch[charIndex].isEnd {
				return true
			} else {
				return false
			}
		}
		currentBranch = &currentBranch[charIndex].children
	}
	return false
}

func getCharIndex(s string) int {
	return int(strings.ToLower(s)[0]) - 97
}

func main() {
	t := Trie{}
	t.Insert("hello")
	fmt.Println(t.Search("hello"))
	fmt.Println(t.Search("hel"))
}