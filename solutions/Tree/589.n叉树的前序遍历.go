/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    var res []int
	var preorder func(x *Node)
	preorder = func(x *Node) {
		if x == nil {
			return
		}
		res = append(res, x.Val)
		for _, child := range x.Children {
			preorder(child)
		}
	}
	preorder(root)
	return res
}
// @lc code=end

