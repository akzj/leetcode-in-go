package best_time_to_buy_and_sell_stock_121

import "math"

func BestTimeToBuyAndSellStock(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	var min = prices[0]
	var maxProfit int
	for i := 1; i < len(prices); i++ {
		min = int(math.Min(float64(min), float64(prices[i])))
		maxProfit = int(math.Max(float64(maxProfit), float64(prices[i]-min)))
	}

	return maxProfit
}
