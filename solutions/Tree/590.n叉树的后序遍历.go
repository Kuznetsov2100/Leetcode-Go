/*
 * @lc app=leetcode.cn id=590 lang=golang
 *
 * [590] N叉树的后序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
	var res []int
	var postorder func(x *Node)
	postorder = func(x *Node) {
		if x == nil {
			return
		}
		for _, child := range x.Children {
			postorder(child)
		}
		res = append(res, x.Val)
	}
	postorder(root)
	return res
}
// @lc code=end

