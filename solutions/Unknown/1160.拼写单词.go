/*
 * @lc app=leetcode.cn id=1160 lang=golang
 *
 * [1160] 拼写单词
 */

// @lc code=start
func countCharacters(words []string, chars string) int {
	arrmap := make([]int, 26)
	for i := range chars {
		arrmap[chars[i]-'a']++
	}

	isValid := func(w string) bool {
		m := make([]int, 26)
		copy(m, arrmap)
		for i := range w {
			if m[w[i]-'a'] == 0 {
				return false
			} else {
				m[w[i]-'a']--
			}
		}
		return true
	}
	var res int
	for _, w := range words {
		if isValid(w) {
			res += len(w)
		}
	}
	return res
}
// @lc code=end

