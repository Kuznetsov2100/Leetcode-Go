/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	// 关键词：前缀和，数据离散化，线段树
	// 题意求数组内任意区间和的范围在[lower, upper]的个数，
	// lower <= S(i, j) <= upper，
	// lower <= prefixSum[j]-prefixSum[i-1] <= upper,
	// lower+prefixSum[i-1] <= prefixSum[j] <= upper+prefixSum[i-1]
	// 数据离散化的原因：prefixSum数组中的数字之间的间隔可能很大，不方便构造线段树，
	// 利用哈希将它映射到一段连续的区间，要注意在映射之前需将allNums数组排序，
	// 否则会破坏线段树中的节点定义。
	var res int
	n := len(nums)
	if n == 0 {
		return res
	}

	var allNums []int
	prefixSum := make([]int, n+1) // 前缀和
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + nums[i]
		allNums = append(allNums, prefixSum[i+1], prefixSum[i]+lower, prefixSum[i]+upper)
	}
	sort.Ints(allNums)
	m := make(map[int]int)
	index := 0
	for _, num := range allNums {
		if _, ok := m[num]; !ok {
			m[num] = index
			index++
		}
	}
	
	root := buildSegmentTree(0, index-1) // 在区间[0, index-1]上构造线段树
	for i := n-1; i >= 0; i-- {
		update(m[prefixSum[i+1]], root) // update操作将prefixSum[i+1]的映射值加入线段树中
		// 通过查询区间[m[lower+prefixSum[i]], m[upper+prefixSum[i]]],
		// 即可知道S(i, i), S(i, i+1), S(i, i+2)....S(i, j)落在[lower, upper]的个数， i <= j
		res += query(m[lower+prefixSum[i]], m[upper+prefixSum[i]], root)
	}
	return res
}

type TreeNode struct {
	start, end, count int
	Left,Right *TreeNode
}

func buildSegmentTree(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{start:start, end: end}
	if start < end {
		node.Left = buildSegmentTree(start, mid)
		node.Right = buildSegmentTree(mid+1, end)
	}
	return node
}

func update(index int, node *TreeNode) {
	if node == nil {
		return
	}
	if index > node.end || index < node.start {
		return
	}
	node.count++
	update(index, node.Left)
	update(index, node.Right)
}

func query(min, max int, node *TreeNode) int {
	if node == nil || min > max {
		return 0
	}
	if node.start == min && node.end == max {
		return node.count
	}
	mid := node.start + (node.end-node.start)/2
	var ans int
	if max <= mid {
		ans = query(min, max, node.Left)
	} else {
		if min <= mid {
			ans = query(min, mid, node.Left) + query(mid+1, max, node.Right)
		} else {
			ans = query(min, max, node.Right)
		}
	}
	return ans
}
// @lc code=end

