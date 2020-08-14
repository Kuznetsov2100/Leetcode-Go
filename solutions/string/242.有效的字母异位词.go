/*
 * @lc app=leetcode.cn id=242 lang=golang
 *
 * [242] 有效的字母异位词
 */

// @lc code=start
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int,26)
	for i := 0; i < len(s); i++ {
		count[s[i]-97]++
		count[t[i]-97]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

