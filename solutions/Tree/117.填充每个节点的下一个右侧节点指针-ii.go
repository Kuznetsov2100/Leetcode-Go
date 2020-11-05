/*
 * @lc app=leetcode.cn id=117 lang=golang
 *
 * [117] 填充每个节点的下一个右侧节点指针 II
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */
func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left != nil {
		root.Left.Next = root.Right
	}
	if root.Next != nil {
		if root.Right != nil {
			root.Right.Next = helper(root.Next)
		} else if root.Left != nil {
			root.Left.Next = helper(root.Next)
		}
	}
	// 先右后左
	connect(root.Right)
	connect(root.Left)
	return root
}

// 递归版本改成迭代版本
func helper(x *Node) *Node {
	for x != nil {
		if x.Left != nil {
			return x.Left
		}
		if x.Right != nil {
			return x.Right
		}
		x = x.Next
	}
	return nil
}
// @lc code=end

