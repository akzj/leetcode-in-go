package reverse_linked_list_206

import (
	. "github.com/akzj/leetcode-in-go/structures"
)

func ReverseList(head *ListNode) *ListNode {
	var reverse func(head *ListNode) *ListNode
	var newHead *ListNode
	reverse = func(node *ListNode) *ListNode {
		if node == nil {
			return nil
		}
		next := reverse(node.Next)
		if next == nil {
			newHead = node
			return node
		}
		next.Next = node
		return node
	}
	reverse(head)
	head.Next = nil
	return newHead
}

func ReverseList2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var pre *ListNode
	var curr = head
	for {
		next := curr.Next
		curr.Next = pre
		pre = curr
		if next == nil {
			return curr
		}
		curr = next
	}
}
