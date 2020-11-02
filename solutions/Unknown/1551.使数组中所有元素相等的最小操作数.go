/*
 * @lc app=leetcode.cn id=1551 lang=golang
 *
 * [1551] 使数组中所有元素相等的最小操作数
 */

// @lc code=start
func minOperations(n int) int {
	// 注意到本题给定的操作并不会使数组中所有元素之和变化，且题目要求让所有元素相等，
	// 那么数组中所有元素的平均值即为最后数组中每一个元素的值。
	// 根据等差数列求和公式可得平均数为n
	// 我们只要统计所有大于平均数的数与n的差值，就能计算出我们操作了多少次。
	var res int
	avg := n
	for i := 0; i < n; i++ {
		num := 2 * i + 1
		if num > avg {
			res += num-avg
		}
	}
	return res
}
// @lc code=end

