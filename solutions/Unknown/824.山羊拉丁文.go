/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

// @lc code=start
func toGoatLatin(S string) string {
	var words []string
	vowels := []byte{'a','e','i','o', 'u','A','E','I', 'O', 'U'}
	isVowel := func(x byte) bool {
		for i := range vowels {
			if vowels[i] == x {
				return true
			}
		}
		return false
	}

	for i, w := range strings.Split(S, " ") {
		if isVowel(w[0]) {
			words = append(words, w + "ma" + strings.Repeat("a", i+1))
		} else {
			words = append(words, w[1:] + w[:1] + "ma" + strings.Repeat("a", i+1))
		}	
	}
	return strings.Join(words, " ")
}
// @lc code=end

