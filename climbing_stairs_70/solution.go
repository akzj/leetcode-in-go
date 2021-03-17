package climbing_stairs_70

import "fmt"

func ClimbingStairs1(stairs int) int {
	return exploreClimbingStair1(stairs, make(map[int]int))
}

func exploreClimbingStair1(stairs int, memo map[int]int) int {
	if stairs == 0 {
		return 1
	}
	if stairs == 1 {
		return 1
	}

	if c, ok := memo[stairs]; ok {
		return c
	}

	c := exploreClimbingStair1(stairs-1, memo) + exploreClimbingStair1(stairs-2, memo)
	fmt.Println(stairs,c)
	memo[stairs] = c
	return c
}
