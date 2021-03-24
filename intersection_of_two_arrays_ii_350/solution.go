package intersection_of_two_arrays_ii_350

func IntersectionOfTwoArrays(first, second []int) []int {
	if len(first) > len(second) {
		second, first = first, second
	}

	nums := make(map[int]int)

	for _, num := range second {
		nums[num]++
	}

	var result []int
	for _, num := range first {
		if val, ok := nums[num]; ok && (val > 0) {
			result = append(result, val)
			nums[num]--
		}
	}
	return result
}
