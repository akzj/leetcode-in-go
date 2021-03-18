package container_with_most_water_11

import "math"

func MaxAreaOfWater(heights []int) int {
	var left = 0
	var right = len(heights) - 1

	var maxArea = 0
	for left < right {
		maxArea = int(math.Max(
			math.Min(float64(heights[left]),
				float64(heights[right]))*float64(right-left),
			float64(maxArea)))
		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
