package house_robber_198

import "math"

func HouseRobber(moneys []int) int {
	if len(moneys) == 0 {
		return 0
	}
	if len(moneys) == 1 {
		return moneys[0]
	}

	dp := make([]int, len(moneys)+1)
	dp[1] = moneys[0]

	for i := 2; i < len(dp); i++ {
		dp[i] = int(math.Max(float64(dp[i-2]+moneys[i-1]), float64(dp[i-1])))
	}

	return dp[len(dp)-1]
}
