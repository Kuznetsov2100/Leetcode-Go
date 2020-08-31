/*
 * @lc app=leetcode.cn id=373 lang=golang
 *
 * [373] 查找和最小的K对数字
 */

// @lc code=start
import "container/heap"
func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	var res [][]int
	len1, len2 := len(nums1), len(nums2)
	if len1 == 0 || len2 == 0 || k < 1 {
		return res
	}
	visited := make([][]bool, len1)
	for i := range visited {
		visited[i] = make([]bool, len2)
	}
	h := &IntMinHeap{&pair{index1:0, index2:0, sum: nums1[0]+nums2[0]}}
	heap.Init(h)
	for k > 0 && h.Len() > 0 {
		p := heap.Pop(h).(*pair)
		i, j := p.index1, p.index2
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		res = append(res, []int{nums1[i], nums2[j]})
		if i < len1-1 {
			heap.Push(h, &pair{index1:i+1, index2:j, sum: nums1[i+1]+nums2[j]})
		}
		if j < len2-1 {
			heap.Push(h, &pair{index1:i, index2:j+1, sum: nums1[i]+nums2[j+1]})
		}
		k--
	}
	return res
}

type pair struct {
	index1 int
	index2 int
	sum int
}
type IntMinHeap []*pair
func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(*pair)) }
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
// @lc code=end

