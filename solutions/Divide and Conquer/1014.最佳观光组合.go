/*
 * @lc app=leetcode.cn id=1014 lang=golang
 *
 * [1014] 最佳观光组合
 */

// @lc code=start
func maxScoreSightseeingPair(A []int) int {
	// premax为(i < j)的A[i]+i的最大值
	// A[i]+A[j]+i-j = A[i]+i + A[j]-j
	// 固定A[j]-j，premax+A[j]-j即为A[i]+A[j]+i-j的最大值
	var res int
	premax := A[0] + 0 
	for j := 1; j < len(A); j++ {
		res = max(res, premax+A[j]-j) // 更新最高分
		premax = max(premax, A[j]+j) //  更新a[i]+i的最大值
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

