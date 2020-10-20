/*
 * @lc app=leetcode.cn id=830 lang=golang
 *
 * [830] 较大分组的位置
 */

// @lc code=start
func largeGroupPositions(s string) [][]int {
	var res [][]int
	i, n := 0, len(s)
	for i < n {
		c, firstIndex, count := s[i], i, 0
		for i < n && c == s[i] {
			i++
			count++
		}
		if count >= 3 {
			res = append(res, []int{firstIndex, i-1})
		}
	}
	return res
}
// @lc code=end

