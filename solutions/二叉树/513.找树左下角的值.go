/*
 * @lc app=leetcode.cn id=513 lang=golang
 *
 * [513] 找树左下角的值
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
// func findBottomLeftValue(root *TreeNode) int {
// 	var queue []*TreeNode
// 	var node *TreeNode	
// 	queue = append(queue, root)
// 	for len(queue) > 0 {
// 		node = queue[0]
// 		queue = queue[1:]
// 		if node.Right != nil {
// 			queue = append(queue, node.Right)
// 		}
// 		if node.Left != nil {
// 			queue = append(queue,node.Left)
// 		}
// 	}
// 	return node.Val
	
// }

func findBottomLeftValue(root *TreeNode) int {
	var res []int
	var queue []*TreeNode
	queue = append(queue, root) // //将根节点放入队列中，然后不断遍历队列
	for len(queue) != 0 {
		var level int
		// n := len(queue),获取当前队列的长度，这个长度相当于当前这一层的节点个数	
		// 将队列中的元素都拿出来(也就是获取这一层的节点)，暂存到level中
		// 如果节点的左/右子节点不为空，放入队列中
		n := len(queue)
		for i := n; i > 0; i-- {
			node := queue[0]
			queue = queue[1:]
			if i == n {
				level = node.Val
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level) // 将这一层的元素加入结果中
	}
	return res[len(res)-1]
}
// @lc code=end

