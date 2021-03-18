package delete_node_in_a_linked_list_237

import "github.com/akzj/leetcode-in-go/structures"

func DeleteNode(head *structures.ListNode, val int) *structures.ListNode {
	var pre = head
	var curr = head
	pre = nil
	for curr != nil {
		if curr.Val == val {
			if pre == nil {
				return head.Next
			}
			pre.Next = curr.Next
			curr.Next = nil
			break
		}
		pre = curr
		curr = curr.Next
	}
	return head
}
