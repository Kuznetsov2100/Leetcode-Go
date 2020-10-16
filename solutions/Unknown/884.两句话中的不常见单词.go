/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(A string, B string) []string {
	var res, arr []string
	m := make(map[string]int)
	arr = strings.Split(A, " ")
	arr = append(arr, strings.Split(B, " ")...)
	for _, word := range arr {
		m[word]++
	}
	for k, v := range m {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
// @lc code=end

