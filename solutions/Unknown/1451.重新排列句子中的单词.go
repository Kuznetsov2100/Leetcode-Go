/*
 * @lc app=leetcode.cn id=1451 lang=golang
 *
 * [1451] 重新排列句子中的单词
 */

// @lc code=start
func arrangeWords(text string) string {
	strs := strings.Split(text, " ")
	strs[0] = strings.ToLower(strs[0])
	sort.SliceStable(strs, func(i, j int) bool {return len(strs[i]) < len(strs[j])})
	strs[0] = strings.Title(strs[0])
	return strings.Join(strs, " ")
}
// @lc code=end

