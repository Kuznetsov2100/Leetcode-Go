/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, S string) []int {
	lines, sum := 1, 0
	for i := range S {
		if sum + widths[S[i]-'a'] <= 100 {
			sum += widths[S[i]-'a']
		} else {
			lines++
			sum = widths[S[i]-'a']
		}
	}
	return []int{lines, sum}
}
// @lc code=end

