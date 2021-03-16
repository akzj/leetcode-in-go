// Package structures provides structures that can be used to create various data structures.
package structures

import "fmt"

// TreeNode is a node in a binary tree data structure.
type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

// ListNode is a node in a singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func PrintList(head *ListNode) {
	for {
		if head == nil {
			fmt.Print("null")
			return
		}
		fmt.Print(head.Val)
		fmt.Print(" -> ")
		head = head.Next
	}
}

func PrintTree(head *TreeNode, leaves int) {
	if head == nil {
		return
	}
	fmt.Println(head.Val, leaves)
	PrintTree(head.Left, leaves+1)
	PrintTree(head.Right, leaves+1)
}

func BuildTree(vals []int) *TreeNode {
	if vals == nil {
		return nil
	}
	var node *TreeNode
	var head *TreeNode
	for _, val := range vals {
		if head == nil {
			node = &TreeNode{
				Val: val,
			}
			head = node
			continue
		}
		node = head
		for {
			if val < node.Val {
				if node.Left != nil {
					node = node.Left
				} else {
					node.Left = &TreeNode{
						Val: val,
					}
					break
				}
			} else if val > node.Val {
				if node.Right != nil {
					node = node.Right
				} else {
					node.Right = &TreeNode{
						Val: val,
					}
					break
				}
			} else {
				break
			}
		}
	}
	return head
}
