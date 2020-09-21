/*
 * @lc app=leetcode.cn id=1190 lang=golang
 *
 * [1190] 反转每对括号间的子串
 */

// @lc code=start
func reverseParentheses(s string) string {
	stack := make([]string, 0, len(s))
	strs := make([]string, 0, len(s))
	for i := range s {
		if s[i] == ')' {
			for len(stack) > 0 && stack[len(stack)-1] != "(" {
				strs = append(strs, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			stack = append(stack, strs...)
			strs = nil
		} else {
			stack = append(stack, string(s[i]))
		}
	}
	return strings.Join(stack, "")
}
// @lc code=end

