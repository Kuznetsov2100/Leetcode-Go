/*
 * @lc app=leetcode.cn id=942 lang=golang
 *
 * [942] 增减字符串匹配
 */

// @lc code=start
func diStringMatch(S string) []int {
	left, right := 0, len(S)
	res := make([]int, 0, len(S)+1)
	for i := range S {
		if S[i] == 'I' {
			res = append(res, left)
			left++
		} else {
			res = append(res, right)
			right--
		}
	}
	res = append(res, left)
	return res
}
// @lc code=end

