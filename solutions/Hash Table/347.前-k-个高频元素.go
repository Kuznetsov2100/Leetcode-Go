/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
import "container/heap"
func topKFrequent(nums []int, k int) []int {
	var res []int
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	h := &MaxHeap{}
	heap.Init(h)
	for number, frequency := range m {
		heap.Push(h, &element{value:number, frequency:frequency})
	}
	for k > 0 && h.Len() > 0 {
		res = append(res, heap.Pop(h).(*element).value)
		k--
	}
	return res
}

type element struct {
	value int
	frequency int
}

type MaxHeap []*element
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i].frequency > h[j].frequency }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(*element)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
// @lc code=end

