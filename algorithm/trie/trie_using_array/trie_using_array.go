package trie_using_array

import "github.com/leetcode/algorithm/trie"

type trieNode struct {
	char byte
	term bool
	next [26]*trieNode
}

type Trie struct {
	root *trieNode
}

func newTrieNode(char byte) *trieNode {
	return &trieNode{
		char: char,
		term: false,
		next: [26]*trieNode{},
	}
}

func Constructor() Trie {
	return Trie{
		root: newTrieNode('*'),
	}
}
func ConstructorInterface() trie.TrieInterface {
	return &Trie{
		root: newTrieNode('*'),
	}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	node := this.root
	for i := 0; i < len(word); i++ {
		char := word[i]
		index := word[i] - 'a'
		if node.next[index] == nil {
			node.next[index] = newTrieNode(char)
		}
		node = node.next[index]
	}
	node.term = true
}

func (this *Trie) find(prefix string) *trieNode {
	node := this.root
	for i := 0; i < len(prefix) && node != nil; i++ {
		node = node.next[prefix[i]-'a']
	}
	return node
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	node := this.find(word)
	return node != nil && node.term
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	node := this.find(prefix)
	return node != nil
}
