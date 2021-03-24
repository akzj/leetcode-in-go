package intersection_of_two_linked_lists_160

import (
	"github.com/austingebauer/go-leetcode/structures"
	"math"
)

func IntersectionNode(first, second *structures.ListNode) *structures.ListNode {

	firstLen, secondLen := listLength(first), listLength(second)
	lenDiff := int(math.Abs(float64(firstLen - secondLen)))

	if lenDiff > 0 {
		if firstLen > secondLen {
			for lenDiff > 0 {
				first = first.Next
				lenDiff--
			}
		} else {
			for lenDiff > 0 {
				second = second.Next
				lenDiff--
			}
		}
	}

	for first != nil && second != nil {
		if first == second {
			return first
		}
		first = first.Next
		second = second.Next
	}
	return nil
}

func listLength(head *structures.ListNode) int {
	var l int
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
