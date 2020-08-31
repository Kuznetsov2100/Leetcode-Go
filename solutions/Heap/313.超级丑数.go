/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
import "container/heap"
func nthSuperUglyNumber(n int, primes []int) int {
	nums := make([]int, n)
	nums[0] = 1
	p := make([]int, len(primes))
	seen := make(map[int]bool)
	h := &IntHeap{}
	heap.Init(h)
	for i := 1; i < n; i++ {
		for j, prime := range primes {
			if newugly := prime*nums[p[j]]; !seen[newugly] {
				heap.Push(h, newugly)
				seen[newugly] = true
			}	
		}
		for h.Len() > 0 && (*h)[0] == nums[i-1] {
			heap.Pop(h)
		}
		minUgly := heap.Pop(h).(int)
		nums[i] = minUgly
		for k, prime := range primes {
			if minUgly == nums[p[k]]*prime {
				p[k]++
			}
		}
	}
	return nums[n-1]
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
// @lc code=end

