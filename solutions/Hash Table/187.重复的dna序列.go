/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) []string {
	var res []string
	n := len(s)
	if n < 10 {
		return res
	}
	m := make(map[string]int)
	for i, j := 0, 9; i < n && j < n; i, j = i+1, j+1 {
		m[s[i:j+1]]++
	}
	for k, v := range m {
		if v > 1 {
			res = append(res, k)
		}
	}
	return res
	
}
// @lc code=end

