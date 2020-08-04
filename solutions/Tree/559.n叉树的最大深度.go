/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	maxmium := -1
	for _, child := range root.Children {
		maxmium = max(maxDepth(child)+1,maxmium)
	}
	return maxmium
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

