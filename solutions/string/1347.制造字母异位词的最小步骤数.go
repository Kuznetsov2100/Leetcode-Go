/*
 * @lc app=leetcode.cn id=1347 lang=golang
 *
 * [1347] 制造字母异位词的最小步骤数
 */

// @lc code=start
func minSteps(s string, t string) int {
	var res int
	count := make([]int,26)
	for i := 0; i < len(s); i++ {
		count[s[i]-97]++
		count[t[i]-97]--
	}
	for i := range count {
		if count[i] > 0 {
			res += count[i]
		}
	}
	return res
}
// @lc code=end

