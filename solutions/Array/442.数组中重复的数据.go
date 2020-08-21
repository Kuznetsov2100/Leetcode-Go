/*
 * @lc app=leetcode.cn id=442 lang=golang
 *
 * [442] 数组中重复的数据
 */

// @lc code=start
func findDuplicates(nums []int) []int {
	var res []int
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for k, v := range m {
		if v == 2 {
			res = append(res, k)
		}
	}
	return res
}
// @lc code=end

