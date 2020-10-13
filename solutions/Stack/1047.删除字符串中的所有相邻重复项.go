/*
 * @lc app=leetcode.cn id=1047 lang=golang
 *
 * [1047] 删除字符串中的所有相邻重复项
 */

// @lc code=start
func removeDuplicates(S string) string {
	var stack []byte
	for i := range S {
		if len(stack) > 0 && stack[len(stack)-1] == S[i] {
			stack = stack[:len(stack)-1]
			continue
		}
		stack = append(stack, S[i])
	}
	return string(stack)
}
// @lc code=end

