/*
 * @lc app=leetcode.cn id=1455 lang=golang
 *
 * [1455] 检查单词是否为句中其他单词的前缀
 */

// @lc code=start
func isPrefixOfWord(sentence string, searchWord string) int {
	for i , s := range strings.Split(sentence, " ") {
		if strings.HasPrefix(s, searchWord) {
			return i+1
		}
	}
	return -1
}
// @lc code=end

