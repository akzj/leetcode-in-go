package combination_sum_iv_377

func CombinationSum4(nums []int, target int) int {
	var output int
	combinationSum4(nums, target, &output)
	return output
}

func combinationSum4(nums []int, target int, output *int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*output += 1
	}
	for _, num := range nums {
		combinationSum4(nums, target-num, output)
	}
}

func CombinationSum42(nums []int, target int) int {

	var dp = make([]int, target+1)
	dp[0] = 1

	for i := 1; i <= target; i++ {

		for _, num := range nums {
			if i < num {
				continue
			}

			dp[i] += dp[i-num]
		}
	}
	return dp[len(dp)-1]
}
