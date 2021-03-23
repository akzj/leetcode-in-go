package insert_interval_57

import (
	"fmt"
	"testing"
)

func TestInsertIntervals(t *testing.T) {

	fmt.Println(InsertIntervals(nil, Interval{1, 2}))

	fmt.Println(InsertIntervals([]Interval{
		{1, 3},
		{6, 9},
	}, Interval{2, 5}))

	fmt.Println(InsertIntervals([]Interval{
	{1,2},{3,5},{6,7},{8,10},{12,16},
	}, Interval{4, 8}))


	fmt.Println(InsertIntervals([]Interval{
		{1,2},{3,5},{6,7},{8,10},{12,16},
	}, Interval{1, 3}))



}

func TestFindIndex(t *testing.T) {
	fmt.Println(findIndex([]Interval{
		{1,2},{3,5},{6,7},{8,10},{12,16},
	},4))
}