package main

import (
	"fmt"
	"reflect"
	"time"

	"github.com/leetcode/algorithm/trie"
	"github.com/leetcode/algorithm/trie/trie_using_array"
	"github.com/leetcode/algorithm/trie/trie_using_map"
)

func testTrieSpeed(t trie.TrieInterface) {
	s := ""
	startAt := time.Now()
	for i := 1; i < 10000; i++ {
		s += "a"
		t.Search(s)
		t.Insert(s)
		t.Search(s)
		t.Search(s + "x")
		t.StartsWith(s)
		t.StartsWith(s + "x")
	}
	fmt.Println(reflect.TypeOf(t), time.Since(startAt))
}

func main() {
	testTrieSpeed(trie_using_array.ConstructorInterface())
	testTrieSpeed(trie_using_map.ConstructorInterface())
}
