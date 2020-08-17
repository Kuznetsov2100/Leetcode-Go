/*
 * @lc app=leetcode.cn id=645 lang=golang
 *
 * [645] 错误的集合
 */

// @lc code=start
func findErrorNums(nums []int) []int {
	res := make([]int, 2)
	m := make([]int,len(nums)+1)
	for i := range nums {
		m[nums[i]]++
	}
	m[0] = 100
	for k, v := range m {
		if v == 0 {
			res[1] = k
			if res[0] != 0 {
				break
			}
		}
		if v == 2 {
			res[0] = k
			if res[1] != 0 {
				break
			}
		}
	}
	return res
}
// @lc code=end

