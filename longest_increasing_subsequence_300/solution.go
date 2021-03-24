package longest_increasing_subsequence_300

import (
	"math"
)

func LongestIncreasingSubsequence2(nums []int) int {

	dp := make([]int, len(nums))
	dp[0] = 1

	var max int
	for i := 0; i < len(nums); i++ {
		var jMax int
		var find bool
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				find = true
				if dp[j] > jMax {
					jMax = dp[j]
				}
			}
		}
		if find {
			jMax++
			dp[i] = jMax
			if jMax > max {
				max = jMax
			}
		}
	}

	return max

}
func LongestIncreasingSubsequence(nums []int) int {
	var max int
	var memo = make(map[int]int)
	for i := range nums {
		n := longestIncreasingSubsequence(nums, i, memo)
		if n > max {
			max = n + 1
		}
	}
	return max
}

func longestIncreasingSubsequence(nums []int, index int, memo map[int]int) int {

	if v, ok := memo[index]; ok {
		return v
	}

	var max = 0
	for i := index + 1; i < len(nums); i++ {
		if nums[index] < nums[i] {
			max = int(math.Max(float64(max), float64(longestIncreasingSubsequence(nums, i, memo)+1)))
		}
	}
	memo[index] = max
	return max
}
