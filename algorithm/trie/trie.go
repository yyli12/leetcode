package trie

type TrieInterface interface {
	Insert(word string)
	Search(word string) bool
	StartsWith(prefix string) bool
}
