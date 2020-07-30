/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) []int {
	var res []int
	var queue []*TreeNode
	
	if root != nil {
		queue = append(queue, root) // //将根节点放入队列中，然后不断遍历队列
	}
	for len(queue) != 0 {
		// n := len(queue),获取当前队列的长度，这个长度相当于当前这一层的节点个数	
		// 将队列中的元素都拿出来(也就是获取这一层的节点)，暂存到level中
		// 如果节点的左/右子节点不为空，放入队列中
		maxValue := -1 << 63 // 初始化为int类型最小值
		for n := len(queue); n > 0; n-- {
			node := queue[0]
			queue = queue[1:]
			maxValue = max(node.Val,maxValue)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, maxValue) // 将这一层中最大的元素加入结果中
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

