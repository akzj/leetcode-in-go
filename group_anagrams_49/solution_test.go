package group_anagrams_49

import (
	"fmt"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {

	for _, group := range GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}) {
		fmt.Println(group)
	}
}
