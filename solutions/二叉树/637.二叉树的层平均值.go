/*
 * @lc app=leetcode.cn id=637 lang=golang
 *
 * [637] 二叉树的层平均值
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
func averageOfLevels(root *TreeNode) []float64 {
	var res []float64
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root) // //将根节点放入队列中，然后不断遍历队列
	}
	for len(queue) != 0 {
		var level []int
		// n := len(queue),获取当前队列的长度，这个长度相当于当前这一层的节点个数	
		// 将队列中的元素都拿出来(也就是获取这一层的节点)，暂存到level中
		// 如果节点的左/右子节点不为空，放入队列中
		for n := len(queue); n > 0; n-- {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		sum := 0
		for i := range level {
			sum += level[i]
		}
		average := float64(sum)/float64(len(level))
		res = append(res, average) // 将这一层的元素的平均值加入结果中
	}
	return res
}
// @lc code=end

