/*
 * @lc app=leetcode.cn id=453 lang=golang
 *
 * [453] 最小移动次数使数组元素相等
 */

// @lc code=start
func minMoves(nums []int) int {
	// n-1个元素增加1,这相当于一个数减去1，对数组求和， 同时找到最小值，
	// 计算将数组中的数全部减到最小值需要多少次减法
	var sum int
	minimum := math.MaxInt64
	for _, num := range nums {
		sum += num
		if num < minimum {
			minimum = num
		}
	}
	return sum-minimum*len(nums)
}
// @lc code=end

