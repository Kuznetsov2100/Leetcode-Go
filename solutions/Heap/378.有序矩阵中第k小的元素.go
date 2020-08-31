/*
 * @lc app=leetcode.cn id=378 lang=golang
 *
 * [378] 有序矩阵中第K小的元素
 */

// @lc code=start
import "container/heap"
func kthSmallest(matrix [][]int, k int) int {
	var res int
	n := len(matrix)
	h := &IntMinHeap{&foo{value:matrix[0][0], row:0, col:0 }}
	heap.Init(h)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}
	for k > 0 && h.Len() > 0 {
		element := heap.Pop(h).(*foo)
		i, j := element.row, element.col
		if visited[i][j] {
			continue
		}
		visited[i][j] = true
		res = element.value
		if i < n-1 {
			heap.Push(h, &foo{value:matrix[i+1][j], row:i+1, col:j})
		}
		if j < n-1 {
			heap.Push(h, &foo{value:matrix[i][j+1], row:i, col:j+1})
		}
		k--
	}
	return res
}

type foo struct {
	value int
	row int
	col int
}
type IntMinHeap []*foo
func (h IntMinHeap) Len() int           { return len(h) }
func (h IntMinHeap) Less(i, j int) bool { return h[i].value < h[j].value }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntMinHeap) Push(x interface{}) { *h = append(*h, x.(*foo)) }
func (h *IntMinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
// @lc code=end

