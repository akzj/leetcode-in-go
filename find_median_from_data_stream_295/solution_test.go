package find_median_from_data_stream_295

import (
	"fmt"
	"testing"
)

func TestNewMedianQueue(t *testing.T) {

	queue := NewMedianQueue()

	queue.AddNum(1)
	queue.AddNum(2)

	fmt.Println(queue.findMedian())

	queue.AddNum(3)
	fmt.Println(queue.findMedian())

	queue.AddNum(4)
	fmt.Println(queue.findMedian())


	queue.AddNum(5)
	fmt.Println(queue.findMedian())
}
