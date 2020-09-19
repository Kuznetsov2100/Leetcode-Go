/*
 * @lc app=leetcode.cn id=32 lang=golang
 *
 * [32] 最长有效括号
 */

// @lc code=start
// 始终保持栈底元素为当前已经遍历过的元素中「最后一个没有被匹配的右括号的下标」,
// 栈中其它元素为左括号的下标
// 对于遇到的每个'(', 我们将它的下标放入栈中
// 对于遇到的每个')', 我们先弹出栈顶元素表示匹配了当前右括号：
// 如果栈为空，说明当前的右括号为没有被匹配的右括号，我们将其下标放入栈中来更新我们之前提到的「最后一个没有被匹配的右括号的下标」
// 如果栈不为空，当前右括号的下标减去栈顶元素即为「以该右括号为结尾的最长有效括号的长度」
// 需要注意的是，如果一开始栈为空，第一个字符为左括号的时候我们会将其放入栈中，
// 这样就不满足提及的「最后一个没有被匹配的右括号的下标」，为了保持统一，我们在一开始的时候往栈中放入值为-1的元素
func longestValidParentheses(s string) int {
	var res int
	stack := []int{-1} // 初始化栈
	for i := range s {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i) // 更新最后一个没有被匹配的右括号的下标
			} else {
				res = max(res, i-stack[len(stack)-1])
			}
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

