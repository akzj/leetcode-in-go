package best_time_to_buy_and_sell_stock_121

import (
	"fmt"
	"testing"
)

func TestBestTimeToBuyAndSellStock(t *testing.T) {
	fmt.Println(BestTimeToBuyAndSellStock([]int{
		7, 1, 5, 3, 6, 4,
	}))

	fmt.Println(BestTimeToBuyAndSellStock([]int{7, 6, 4, 3, 1}))
}
