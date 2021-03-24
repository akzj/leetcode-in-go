package linked_list_cycle_141

import "github.com/austingebauer/go-leetcode/structures"

func LinkedListCycle(node *structures.ListNode) int {
	fast := node
	slow := node

	for fast != nil && slow != nil && fast != slow {
		if fast.Next != nil {
			fast = fast.Next.Next
		}
		slow = slow.Next
	}
	if fast == nil || slow == nil {
		return -1
	}

	var n int

	for node != slow {
		node = node.Next
		slow = slow.Next
		n++
	}
	return n
}
