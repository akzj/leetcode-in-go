package insert_interval_57

import (
	"fmt"
	"math"
)

type Interval struct {
	Start int
	End   int
}

func (i Interval) String() string {
	return fmt.Sprintf("[%d %d]", i.Start, i.End)
}

func InsertIntervals(intervals []Interval, inter Interval) []Interval {
	if len(intervals) == 0 {
		return append(intervals, inter)
	}

	var mergeIntervals []Interval
	for i := 0; i < len(intervals); i++ {
		if intervals[i].End < inter.Start || intervals[i].Start > inter.End {
			mergeIntervals = append(mergeIntervals, intervals[i])
			continue
		}

		var temp = inter
		for ; i < len(intervals); i++ {
			temp.Start = int(math.Min(float64(temp.Start), float64(intervals[i].Start)))
			if temp.End > intervals[i].End {
				continue
			}
			if temp.End < intervals[i].Start {
				i--
				break
			}
			temp.End = intervals[i].End
			break
		}
		mergeIntervals = append(mergeIntervals, temp)
	}
	return mergeIntervals
}

func findIndex(intervals []Interval, start int) int {

	left := 0
	right := len(intervals)
	for left < right {
		mid := (left + right) / 2
		if start > intervals[mid].End {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return left
}
