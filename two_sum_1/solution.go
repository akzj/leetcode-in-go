package two_sum_1

import "sort"

// twoSum does not assume that nums is sorted.
func twoSum(nums []int, target int) []int {
	numsLen := len(nums)
	if numsLen < 2 {
		return []int{}
	}

	numToIndex := make(map[int]int)
	for idx, num := range nums {
		remain := target - num
		idx2,ok := numToIndex[remain]
		if ok{
			return []int{idx2,idx}
		}
		numToIndex[num] = idx
	}

	return []int{}
}

// twoSumSortedInput assumes that nums is sorted.
func twoSumSortedInput(nums []int, target int) []int {
	sort.Ints(nums)

	numsLen := len(nums)
	if numsLen < 2 {
		return []int{}
	}

	front:=0
	end := len(nums)-1

	for front != end{
		sum := nums[front]+nums[end]
		if sum == target{
			return []int{front,end}
		}else if sum < target{
			front ++
		}else {
			end--
		}
	}
	return nil
}
