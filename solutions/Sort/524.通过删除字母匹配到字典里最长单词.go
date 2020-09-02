/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
import "container/heap"
func findLongestWord(s string, d []string) string {
	var candidate []string
	lens := len(s)
	for i := range d {
		length := len(d[i])
		if length > lens {
			continue
		}
		if isSubsequence(d[i], s, length, lens) {
			candidate = append(candidate, d[i])
		}
	}
	if len(candidate) == 0 {
		return ""
	}
	h := make(MaxHeap, len(candidate))
	copy(h,candidate)
	heap.Init(&h)
	return heap.Pop(&h).(string)
}

func isSubsequence(s string, t string, lens, lent int) bool {
	indexS, indexT := 0, 0
	for indexS < lens && indexT < lent {
		if s[indexS] == t[indexT] {
			indexS++
		}
		indexT++
	}
	return indexS == lens
}

type MaxHeap []string
func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool {
	if m, n := len(h[i]),len(h[j]); m > n {
		return true
	} else if m == n {
		if h[i] < h[j] {
			return true
		}
	}
	return false
}
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(string)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
// @lc code=end

