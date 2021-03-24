package longest_consecutive_sequence_128

import "fmt"

func LongestConsecutiveSequence(nums []int) int {

	indexes := make(map[int]int)

	for _, n := range nums {
		indexes[n] = n
	}

	for n := range indexes {
		if _, ok := indexes[n+1]; ok {
			indexes[n] = n + 1
		}
	}

	fmt.Println(indexes)

	var max int
	for k, v := range indexes {
		var n int
		for k != v {
			n++
			next, ok := indexes[v] //next
			if ok == false {
				break
			}
			k = v
			v = next
		}

		if n > max {
			max = n + 1
		}
	}
	return max
}
