/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
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
import "strconv"
func sumRootToLeaf(root *TreeNode) int {
	var sum int
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {	
			if root == nil {
				return
			}	
			path += strconv.Itoa(root.Val)
			if root.Left == nil && root.Right == nil {
					num, _ := strconv.ParseInt(path,2,0) // 将从根节点到叶子节点的路径字符串按二进制解析
					sum += int(num)
					return
			}
			dfs(root.Left, path)
			dfs(root.Right, path)
			path = path[:len(path)-1]		
	}
	dfs(root, "")
	return sum
}
// @lc code=end

