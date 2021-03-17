package coin_change_322

import (
	"math"
)

func CoinChange2(coins []int, amount int) int {
	var dp = make([]int, amount+1)
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for a := 1; a <= amount; a++ {
		for _, coin := range coins {
			if coin > a {
				continue
			}
			dp[a] = int(math.Min(float64(dp[a]), float64(dp[a-coin]+1)))
		}
	}

	if dp[amount] == amount+1 {
		return -1
	}
	return dp[amount]
}

var dp = make(map[int]int)

func CoinChange(coins []int, amount int) int {
	var ok = false
	for _, coin := range coins {
		if coin == amount {
			return 1
		}
		if coin < amount {
			ok = true
		}
	}

	if ok == false {
		return math.MinInt32
	}

	if c, ok := dp[amount]; ok {
		return c
	}

	var minCount = math.MaxInt32
	for _, coin := range coins {
		if amount < coin {
			continue
		}
		count := CoinChange(coins, amount-coin)
		if count == math.MinInt32 {
			continue
		}
		count += 1
		if count < minCount {
			minCount = count
		}
	}
	if minCount == math.MaxInt32 {
		return math.MinInt32
	}

	dp[amount] = minCount

	return minCount
}
