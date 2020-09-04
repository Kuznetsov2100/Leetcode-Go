/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	// 排列组合知识: C(9,1)*A(9,n-1)+C(9,1)*A(9,n-2)....+C(9,1)*A(9,1) + 10
	res := 1
	for i := 0; i < n; i++ {
		res += 9*factorial(9)/factorial(9-i)
	}
	return res
}

func factorial(n int) int {
	res := 1
	for i :=2; i <= n; i++ {
		res = res*i
	}
	return res
}

// 动态规划
// func countNumbersWithUniqueDigits(n int) int {
// 	if n == 0 {
// 		return 1
// 	}
// 	if n == 1 {
// 		return 10
// 	}
// 	dp := make([]int, n+1)
//  dp[0] = 1
// 	dp[1] = 9
// 	res := dp[0] + dp[1]
// 	for i := 2; i <= n; i++ {
// 		dp[i] = dp[i-1]*(10-(i-1))
// 		res += dp[i]
// 	}
// 	return res
// }
// @lc code=end

