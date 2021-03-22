package find_n_unique_integers_sum_up_to_zero_1304

func FindUniqueSumZero(n int) []int {
	var a = -1
	var b = 1
	var nums []int
	for len(nums) <= n-2 {
		nums = append(nums, a, b)
		a--
		b++
	}
	if len(nums) < n {
		nums = append(nums, 0)
	}
	return nums
}
