package add_two_numbers_2

import (
	"fmt"
	"github.com/akzj/leetcode-in-go/structures"
)

func AddTwoNumbers(first *structures.ListNode, second *structures.ListNode) *structures.ListNode {
	var sumList = &structures.ListNode{}
	var result = sumList

	var flag int
	for first != nil || second != nil {
		var sum int
		if first != nil {
			sum += first.Val
			first = first.Next
		}
		if second != nil {
			sum += second.Val
			second = second.Next
		}
		fmt.Println(sum)
		sum += flag
		flag = sum / 10
		sum %= 10

		sumList.Val = sum
		if first != nil || second != nil {
			sumList.Next = &structures.ListNode{}
			sumList = sumList.Next
		}
	}
	return result
}
