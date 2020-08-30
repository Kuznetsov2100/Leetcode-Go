/*
 * @lc app=leetcode.cn id=264 lang=golang
 *
 * [264] 丑数 II
 */

// @lc code=start
// 动态规划-三指针
func nthUglyNumber(n int) int {
	nums := make([]int, n)
	nums[0] = 1
	p2, p3, p5 := 0, 0, 0
	for i := 1; i < n; i++ {
		minUgly := min(nums[p2]*2, min(nums[p3]*3, nums[p5]*5))
		nums[i] = minUgly
		if minUgly == nums[p2]*2 {
			p2++
		}
		if minUgly == nums[p3]*3 {
			p3++
		}
		if minUgly == nums[p5]*5 {
			p5++
		}
	}
	return nums[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// 最小堆
// import "container/heap"
// func nthUglyNumber(n int) int {
// 	uglynum := -1
// 	h := &IntHeap{1}
// 	heap.Init(h)
// 	for i := 0; i < n; i++ {
// 		for h.Len() > 0 && (*h)[0] == uglynum {
// 			heap.Pop(h)
// 		}
// 		uglynum = heap.Pop(h).(int)
// 		heap.Push(h, 2*uglynum)
// 		heap.Push(h, 3*uglynum)
// 		heap.Push(h, 5*uglynum)
// 	}
// 	return uglynum
// }
// type IntHeap []int

// func (h IntHeap) Len() int           { return len(h) }
// func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
// func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// func (h *IntHeap) Push(x interface{}) {
// 	// Push and Pop use pointer receivers because they modify the slice's length,
// 	// not just its contents.
// 	*h = append(*h, x.(int))
// }

// func (h *IntHeap) Pop() interface{} {
// 	old := *h
// 	n := len(old)
// 	x := old[n-1]
// 	*h = old[0 : n-1]
// 	return x
// }
// @lc code=end

