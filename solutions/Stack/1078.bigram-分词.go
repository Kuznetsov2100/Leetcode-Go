/*
 * @lc app=leetcode.cn id=1078 lang=golang
 *
 * [1078] Bigram 分词
 */

// @lc code=start
func findOcurrences(text string, first string, second string) []string {
	var res []string
	texts := strings.Split(text, " ")
	for i := 0; i < len(texts); i++ {
		if i+2 < len(texts) && texts[i] == first && texts[i+1] == second { 
			res = append(res, texts[i+2])
			i++
		}
	}
	return res
}
// @lc code=end

