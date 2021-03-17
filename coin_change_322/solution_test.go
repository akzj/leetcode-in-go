package coin_change_322

import (
	"fmt"
	"testing"
)

func TestCoinChange(t *testing.T) {
	fmt.Println(CoinChange([]int{1, 2, 5}, 11))

	fmt.Println(CoinChange2([]int{186, 419, 83, 408}, 6249))
}
