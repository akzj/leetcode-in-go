package implement_trie_208

import (
	"fmt"
	"testing"
)

func TestNewTire(t *testing.T) {
	var trie = NewTire()

	trie.Insert("apple")
	fmt.Println(trie.Search("apple"))   // returns true
	fmt.Println(trie.Search("app"))     // returns false
	fmt.Println(trie.startsWith("app")) // returns true
	trie.Insert("app")
	fmt.Println(trie.Search("app")) // returns true

}
