/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	n := len(nums)
	dp_max := make([]int, n)
	dp_min := make([]int, n)
	res := nums[0]
	dp_min[0], dp_max[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		dp_max[i] = max(max(dp_max[i-1]*nums[i], dp_min[i-1]*nums[i]), nums[i])
		dp_min[i] = min(min(dp_max[i-1]*nums[i], dp_min[i-1]*nums[i]), nums[i])
		res = max(res, dp_max[i])
	}
	return res
	
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

