package contains_duplicate_217

func ContainDuplicate(nums []int) bool {
	var m = make(map[int]int)
	for _, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = num
	}
	return false
}
