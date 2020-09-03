/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) int {
	res := nums[0]
	preMin, preMax := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		tmp := preMax
		preMax = max(max(preMax*nums[i], preMin*nums[i]), nums[i])
		preMin = min(min(tmp*nums[i], preMin*nums[i]), nums[i])
		res = max(res, preMax)
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

