/*
 * @lc app=leetcode.cn id=30 lang=golang
 *
 * [30] 串联所有单词的子串
 */

// @lc code=start
func findSubstring(s string, words []string) []int {
	m, n:= len(words[0]), len(words)
	seen := make(map[string]int)
	for _, word := range words {
		seen[word]++
	}

	check := func(s string) bool {
		existed := make(map[string]int)
		for i := 0; i+m <= m*n; i=i+m {
			if seen[s[i:i+m]] == 0 {
				return false
			} else {
				existed[s[i:i+m]]++
			}
		}
		return reflect.DeepEqual(seen, existed)
	}
	var res []int
	for i := 0; i+m*n <= len(s); i++ {
		if check(s[i:i+m*n]) {
			res = append(res, i)
		}
	}
	return res
}
// @lc code=end

