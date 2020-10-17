/*
 * @lc app=leetcode.cn id=819 lang=golang
 *
 * [819] 最常见的单词
 */

// @lc code=start
func mostCommonWord(paragraph string, banned []string) string {
	var res string
	max := 0
	f := func(r rune) bool {return strings.ContainsRune(" !?',;.", r)}
	m := make(map[string]int)
	ban := make(map[string]bool)
	for i := range banned { ban[banned[i]] = true }
	for _, word := range strings.FieldsFunc(paragraph, f) {
		word = strings.ToLower(word)
		if !ban[word] {
			m[word]++
			if v := m[word]; v > max {
				max = v
				res = word
			}
		}	
	}
	return res
}
// @lc code=end

