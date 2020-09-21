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
	var b bytes.Buffer
	for _, w := range strings.Split(S, " ") {
		b.WriteByte('a')
		if isVowel(w[0]) {
			words = append(words, w + "ma" + b.String())
		} else {
			words = append(words, w[1:] + w[:1] + "ma" + b.String())
		}	
	}
	return strings.Join(words, " ")
}
// @lc code=end

