package reverse_linked_list_206

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
	"testing"
)

func TestReverseList(t *testing.T) {

	var list = &structures.ListNode{}
	var head = list
	for i := 0; i < 10; i++ {
		list.Val = i
		if i < 9 {
			list.Next = &structures.ListNode{}
			list = list.Next
		}
	}

	structures.PrintList(head)
	fmt.Println("")
	head = ReverseList2(head)

	structures.PrintList(head)
	fmt.Println("")
}
