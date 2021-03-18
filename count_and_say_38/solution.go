package count_and_say_38

import (
	"strconv"
)

func CountAndSay(n int) string {
	dp := make([]string, n+1)

	dp[1] = "1"

	for i := 2; i <= n; i++ {
		str := dp[i-1]

		var lastC = str[0]
		var count int
		var say string
		for _, c := range str {
			if uint8(c) == lastC {
				count++
			} else {
				say += strconv.Itoa(count) + string(lastC)
				lastC = uint8(c)
				count = 1
			}
		}
		say += strconv.Itoa(count) + string(lastC)
		dp[i] = say
	}

	return dp[n]
}
