/*
 * @lc app=leetcode.cn id=872 lang=golang
 *
 * [872] 叶子相似的树
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
import "reflect"
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var arr1, arr2 []int
	var dfs func(root *TreeNode , choose int)	
	dfs = func(root *TreeNode, choose int) {
		if root != nil {
			if root.Left == nil && root.Right == nil {
				if choose == 1 {
					arr1 = append(arr1, root.Val)
				} else {
					arr2 = append(arr2, root.Val)
				}
			} else {
				dfs(root.Left, choose)
				dfs(root.Right, choose)
			}
		}
	}
	dfs(root1,1)
	dfs(root2, 2)
	return reflect.DeepEqual(arr1, arr2)
}
// @lc code=end

