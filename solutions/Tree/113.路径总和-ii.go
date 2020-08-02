/*
 * @lc app=leetcode.cn id=113 lang=golang
 *
 * [113] 路径总和 II
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
func pathSum(root *TreeNode, sum int) [][]int {
	var paths [][]int
	var dfs func(root *TreeNode, sum int, path []int)
	dfs = func(root *TreeNode, sum int, path []int) {	
			if root == nil {
				return
			}	
			sum = sum - root.Val 
			path = append(path, root.Val)
			if root.Left == nil && root.Right == nil && sum == 0 {
					tmp := make([]int, len(path))
					copy(tmp, path)
					paths = append(paths, tmp)	// 将合法路径的副本加入到结果中，防止后续遍历操作path的时候更改了结果中的数据	
					return
			}
			dfs(root.Left, sum, path)
			dfs(root.Right, sum, path)
			path = path[:len(path)-1]		
	}
	dfs(root, sum, []int{})
	return paths
}
// @lc code=end

