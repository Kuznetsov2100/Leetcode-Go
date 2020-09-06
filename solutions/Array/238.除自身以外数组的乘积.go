/*
 * @lc app=leetcode.cn id=238 lang=golang
 *
 * [238] 除自身以外数组的乘积
 */

// @lc code=start
func productExceptSelf(nums []int) []int {
	n := len(nums)
	front := make([]int, n+1) // 前缀积
	back  := make([]int, n+1) // 后缀积
	front[0] = 1
	back[n]  = 1
	for i, j := 0, n-1; i < n && j >= 0; i, j = i+1, j-1 {
		front[i+1] = front[i] * nums[i]
		back[j] = back[j+1] * nums[j]
	}
	var res []int
	for i := 0; i < n; i++ {
		// output[i]为nums[i]的前缀积的后缀积的乘积
		res = append(res, front[i] * back[i+1])
	}
	return res
}
// @lc code=end

