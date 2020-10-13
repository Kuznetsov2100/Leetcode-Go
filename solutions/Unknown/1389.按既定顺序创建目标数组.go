/*
 * @lc app=leetcode.cn id=1389 lang=golang
 *
 * [1389] 按既定顺序创建目标数组
 */

// @lc code=start
func createTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums))
	for i, ix := range index {
		copy(res[ix+1:], res[ix:])
		res[ix] = nums[i]
	}
	return res
}
// @lc code=end

