/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 */

// @lc code=start
import "container/heap"
// 思路：
// 维护一个最大堆，一个最小堆
// 最大堆存储数据流值较小的一半元素，最小堆存储数据流值较大的一半元素
// case 1: 数据流总元素个数为奇数，最大堆比最小堆多一个元素，中位数为最大堆的堆顶元素
// case 2: 数据流总元素个数为偶数，最大堆与最小堆元素一样多，中位数为最大堆和最小堆的对顶元素之和的一半
// 如何保持最大堆和最小堆的元素处于平衡状态(两个堆的元素个数之差最多为1)
// 步骤：新元素添加到最大堆中，将最大堆的最大值移到最小堆中，
//      如果最大堆的元素少与最小堆，将最小堆的最小值移到最大堆中
type MedianFinder struct {
	minheap *IntMinHeap
	maxheap *IntMaxHeap
}


/** initialize your data structure here. */
func Constructor() MedianFinder {
	minh, maxh := &IntMinHeap{}, &IntMaxHeap{}
	heap.Init(minh)
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
	*h = old[:n-1]
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
	*h = old[:n-1]
	return x
}
/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

