package distribute_candies_575

import "math"

func DistributeCandies(Candies []int) int {

	var kindes = map[int]bool{}
	for _, kind := range Candies {
		kindes[kind] = true
	}

	return int(math.Min(float64(len(kindes)), float64(len(Candies)/2)))
}
