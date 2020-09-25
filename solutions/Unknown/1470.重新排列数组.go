/*
 * @lc app=leetcode.cn id=1470 lang=golang
 *
 * [1470] 重新排列数组
 */

// @lc code=start
func shuffle(nums []int, n int) []int {
	res := make([]int, 0, n*2)
	for i := 0; i < n; i++ {
		res = append(res,nums[i], nums[i+n])
	}
	return res
}
// @lc code=end

