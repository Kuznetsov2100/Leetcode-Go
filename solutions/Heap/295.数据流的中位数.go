/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
import "container/heap"
type MedianFinder struct {
	minheap *IntMinHeap
	maxheap *IntMaxHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	minh := &IntMinHeap{}
	heap.Init(minh)
	maxh := &IntMaxHeap{}
	heap.Init(maxh)
	return MedianFinder{minheap:minh, maxheap:maxh}
}


func (this *MedianFinder) AddNum(num int)  {
	heap.Push(this.maxheap, num)
	heap.Push(this.minheap, heap.Pop(this.maxheap).(int))
	if this.maxheap.Len() < this.minheap.Len() {
		heap.Push(this.maxheap, heap.Pop(this.minheap).(int))
	}	
}


func (this *MedianFinder) FindMedian() float64 {
	if this.maxheap.Len() > this.minheap.Len() {
		return float64((*this.maxheap)[0])
	}
	return float64((*this.minheap)[0]+(*this.maxheap)[0])*0.5
}


type IntMinHeap []int
func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type IntMaxHeap []int
func (h IntMaxHeap) Len() int           { return len(h) }
func (h IntMaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntMaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

