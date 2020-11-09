/*
 * @lc app=leetcode.cn id=315 lang=golang
 *
 * [315] 计算右侧小于当前元素的个数
 */

// @lc code=start
func countSmaller(nums []int) []int {
	// 线段树
	res := make([]int, len(nums))
	minimum := math.MaxInt64
	maximum := math.MinInt64
	for _, num := range nums {
		if num > maximum {
			maximum = num
		}
		if num < minimum {
			minimum = num
		}
	}
	root := buildSegmentTree(minimum, maximum)
	for i := len(nums)-1; i >= 0; i-- {
		res[i] = query(minimum, nums[i]-1, root)
		update(nums[i], root)
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

