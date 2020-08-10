/*
 * @lc app=leetcode.cn id=662 lang=golang
 *
 * [662] 二叉树最大宽度
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//  每个节点一个position值，如果我们走向左子树，那么position -> position * 2，
//  如果我们走向右子树，那么 position -> positon * 2 + 1。
//  当我们在看同一层深度的位置值 L 和 R 的时候，宽度就是 R - L + 1
func widthOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 1
	queue := []*node{{0, root}}
	for len(queue) > 0 {
		var nextQueue []*node
		res = max(res, queue[len(queue)-1].pos-queue[0].pos+1)
		for _, n := range queue {
			if n.Left != nil {
				nextQueue = append(nextQueue, &node{n.pos*2, n.Left})
			}
			if n.Right != nil {
				nextQueue = append(nextQueue, &node{n.pos*2+1, n.Right})
			}
		}
		queue = nextQueue
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

type node struct {
	pos int
	*TreeNode // 内嵌结构体，node继承了TreeNode的字段
}
// @lc code=end

