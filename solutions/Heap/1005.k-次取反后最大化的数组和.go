/*
 * @lc app=leetcode.cn id=1005 lang=golang
 *
 * [1005] K 次取反后最大化的数组和
 */

// @lc code=start
import "container/heap"
func largestSumAfterKNegations(A []int, K int) int {
	// 维护一个最小堆，每次pop最小元素，将该元素取反后push回去，执行K次后结束，最后遍历堆中元素求和。
	// 当弹出的数 >= 0， 根据K的奇偶性，做出相应选择后可提前结束。
	var res int
	n := len(A)
	h := make(IntMinHeap, n)
	copy(h, A)
	heap.Init(&h)
	for K > 0 {
		num := heap.Pop(&h).(int)
		if num < 0 {
			heap.Push(&h, -num)
		} else {
			if K%2 == 0 {
				heap.Push(&h, num)
			} else {
				heap.Push(&h, -num)
			}
			break
		}
		K--
	}

	for i := 0; i < n; i++ {
		res += h[i]
	}
	return res
}

type IntMinHeap []int

func (h IntMinHeap) Len() int            { return len(h) }
func (h IntMinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
// @lc code=end

