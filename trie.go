package main

import (
	"strings"
)

type TrieNode struct {
	children [26]*TrieNode
	isEnd    bool
}

type Trie struct {
	root [26]*TrieNode
}

func (t *Trie) Insert(s string) {
	sLen := len(s)
	currentBranch := &t.root
	for i := 0; i < sLen; i++ {
		charIndex := getCharIndex(s[i : i+1])
		if currentBranch[charIndex] == nil {
			currentBranch[charIndex] = &TrieNode{}
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
