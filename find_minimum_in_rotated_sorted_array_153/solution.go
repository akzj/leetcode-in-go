package find_minimum_in_rotated_sorted_array_153

import "fmt"

func FindMinimum(nums []int) int {
	left := 0
	right := len(nums) - 1
	if nums[left] < nums[right] {
		return nums[0]
	}
	fmt.Println(nums)
	for left < right {
		med := left + (right-left)/2
		if nums[left] < nums[med] {
			left = med
		} else {
			right = med
		}
		fmt.Println(nums[left : right+1])
	}

	return nums[left+1]
}
