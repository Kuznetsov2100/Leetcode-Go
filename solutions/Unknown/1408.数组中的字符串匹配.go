/*
 * @lc app=leetcode.cn id=1408 lang=golang
 *
 * [1408] 数组中的字符串匹配
 */

// @lc code=start
func stringMatching(words []string) []string {
	sort.Slice(words, func(i, j int) bool {return len(words[i]) < len(words[j])})
	var res []string
	m := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		for j := i+1; j < len(words); j++ {
			if m[words[i]] {
				continue
			}
			if strings.Contains(words[j], words[i]) {
				res = append(res, words[i])
				m[words[i]] = true
			}
		}
	}
	return res
}
// @lc code=end

