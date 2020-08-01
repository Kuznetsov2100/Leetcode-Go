/*
 * @lc app=leetcode.cn id=96 lang=golang
 *
 * [96] 不同的二叉搜索树
 */

// @lc code=start
func numTrees(n int) int {
	// dp[i]的含义： 节点数为i能构建出的BST种数
	dp := make([]int, n+1)
	// base case
	dp[0] = 1 // 节点数为0，形成空树
	dp[1] = 1 // 节点数为1， 只有单个节点的BST
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			// 左子树用j个节点，右子树用i-j-1个节点，
			// 能构建出dp[j] * dp[i-j-1]种不同的BST，j 从 0 到 i-1
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}
// @lc code=end