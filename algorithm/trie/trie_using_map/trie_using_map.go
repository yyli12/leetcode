package trie_using_map

import "github.com/leetcode/algorithm/trie"

type trieNode struct {
	char byte
	term bool
	next map[byte]*trieNode
}

type Trie struct {
	root *trieNode
}

func newTrieNode(char byte) *trieNode {
	return &trieNode{
		char: char,
		term: false,
		next: map[byte]*trieNode{},
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
		if _, ok := node.next[char]; !ok {
			node.next[char] = newTrieNode(char)
		}
		node = node.next[char]
	}
	node.term = true
}

func (this *Trie) find(prefix string) *trieNode {
	node := this.root
	for i := 0; i < len(prefix) && node != nil; i++ {
		node = node.next[prefix[i]]
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

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
