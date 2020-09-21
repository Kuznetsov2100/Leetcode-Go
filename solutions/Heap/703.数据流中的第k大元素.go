/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第K大元素
 */

// @lc code=start
// 维护一个最小堆，，堆中储存了前k大的元素，调用Add方法来初始化堆
// Add方法具体实现：
// 如果堆长度小于k,直接push新元素
// 当堆中恰好有k个元素时，h[0]即为第k大的元素，假如新元素比h[0]大，则pop h[0]，再push 新元素
import "container/heap"
type KthLargest struct {
	h *IntMinHeap
	k int
}


func Constructor(k int, nums []int) KthLargest {
	h := &IntMinHeap{}
	heap.Init(h)
	kth := KthLargest{h:h, k:k}
	for _, num := range nums {
		kth.Add(num)
	}
	return kth
}


func (this *KthLargest) Add(val int) int {
	if this.h.Len() < this.k {
		heap.Push(this.h, val)
	} else if (*this.h)[0] < val {
		heap.Pop(this.h)
		heap.Push(this.h, val)
	}
	return (*this.h)[0]
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

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

