/*
 * @lc app=leetcode.cn id=1161 lang=golang
 *
 * [1161] 最大层内元素和
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
func maxLevelSum(root *TreeNode) int {
	var arr [][]int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root) // //将根节点放入队列中，然后不断遍历队列
	}
	index := 0
	for len(queue) != 0 {
		index++
		var levelsum int
		// n := len(queue),获取当前队列的长度，这个长度相当于当前这一层的节点个数	
		// 将队列中的元素都拿出来(也就是获取这一层的节点)，暂存到level中
		// 如果节点的左/右子节点不为空，放入队列中
		for n := len(queue); n > 0; n-- {
			node := queue[0]
			queue = queue[1:]
			levelsum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		arr = append(arr, []int{index, levelsum}) // 将这一层的元素加入结果中
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i][1] > arr[j][1] {
			return true
		} else if arr[i][1] == arr[j][1] {
			if arr[i][0] < arr[j][0] {
				return true
			}
		}
		return false
	})
	return arr[0][0]
}
// @lc code=end

