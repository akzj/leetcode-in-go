package find_median_from_data_stream_295

import "container/heap"

type smallHeap []int

func (e *smallHeap) Len() int {
	return len(*e)
}

func (e *smallHeap) Less(i, j int) bool {
	return (*e)[i] < (*e)[j]
}

func (e *smallHeap) Swap(i, j int) {
	(*e)[i], (*e)[j] = (*e)[j], (*e)[i]
}

func (e *smallHeap) Push(x interface{}) {
	*e = append(*e, x.(int))
}

func (e *smallHeap) Pop() interface{} {
	obj := (*e)[len(*e)-1]
	*e = (*e)[:len(*e)-1]
	return obj
}

type bigHeap []int

func (b *bigHeap) Len() int {
	return len(*b)
}

func (b *bigHeap) Less(i, j int) bool {
	return (*b)[i] > (*b)[j]
}

func (b *bigHeap) Swap(i, j int) {
	(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
}

func (b *bigHeap) Push(x interface{}) {
	*b = append(*b, x.(int))
}

func (b *bigHeap) Pop() interface{} {
	var obj = (*b)[len(*b)-1]
	*b = (*b)[:len(*b)-1]
	return obj
}

type MedianQueue struct {
	smallHeap *smallHeap
	bigHeap   *bigHeap
}

func (m *MedianQueue) AddNum(val int) {
	heap.Push(m.smallHeap, val)
	if m.smallHeap.Len() > m.bigHeap.Len() {
		heap.Push(m.bigHeap, heap.Pop(m.smallHeap).(int))
	}
}

func (m *MedianQueue) findMedian() float64 {
	if m.smallHeap.Len() == m.bigHeap.Len() {
		return float64((*m.smallHeap)[0]+(*m.bigHeap)[0]) / 2
	} else {
		return float64((*m.bigHeap)[0])
	}
}

func NewMedianQueue() *MedianQueue {
	return &MedianQueue{
		smallHeap: &smallHeap{},
		bigHeap:   &bigHeap{},
	}
}
