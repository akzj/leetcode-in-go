package house_robber_ii_213

import "math"

func HouseRobberII(moneys []int) int {
	if len(moneys) == 0 {
		return 0
	}

	if len(moneys) == 1 {
		return moneys[0]
	}
	return int(math.Max(float64(houseRobber(moneys[1:])), float64(houseRobber(moneys[:len(moneys)-1]))))
}

func houseRobber(moneys []int) int {
	dp := make([]int, len(moneys)+1)
	dp[1] = moneys[0]
	for i := 2; i <= len(moneys); i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+moneys[i-1]), float64(dp[i])))
	}
	return dp[len(dp)-1]
}
