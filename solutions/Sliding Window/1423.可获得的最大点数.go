/*
 * @lc app=leetcode.cn id=1423 lang=golang
 *
 * [1423] 可获得的最大点数
 */

// @lc code=start
func maxScore(cardPoints []int, k int) int {
	// 动态规划-前缀和   
	// 假设左边取x个, 则右边区k-x个， 枚举得到sum(x)+sum(k-x)的最大值
	var res int
	n := len(cardPoints)
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] += dp[i-1] + cardPoints[i-1]
	}

	for i := 0; i <= k; i++ {
		res = max(res, dp[i] + (dp[n] - dp[n-(k-i)]))
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// func maxScore(cardPoints []int, k int) int {
// 	// 滑动窗口， 求出长度为len(cardPoints)-k的子数组和的最小值
// 	// 剩下的k个元素的和即为最大点数
// 	n := len(cardPoints)
// 	windowSum, totalSum := 0, 0
// 	for i := 0; i < n; i++ {
// 		totalSum += cardPoints[i]
// 		if i < n-k {
// 			windowSum += cardPoints[i]
// 		}
// 	}
// 	minWindowSum := windowSum
// 	for i := n-k; i < n; i++ {
// 		windowSum += cardPoints[i] // 窗口右移一位
// 		windowSum -= cardPoints[i-n+k] // 为保持窗口大小不变，删除窗口最左侧元素
// 		if windowSum < minWindowSum {
// 			minWindowSum = windowSum
// 		}
// 	}
// 	return totalSum-minWindowSum
// }
// @lc code=end

