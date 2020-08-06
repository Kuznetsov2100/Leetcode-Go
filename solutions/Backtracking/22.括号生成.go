/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

// @lc code=start
func generateParenthesis(n int) []string {
	// 1. 一个合法括号组合的左括号数量等于右括号数量。
    // 2. 对于一个合法的括号字符串组合p，对于任何0 <= i < len(p)都有：子串p[0..i]中左括号的数量都大于或等于右括号的数量。
	var res []string
	var backtrack func(left, right int, track string)
	// left表示剩余可用的左括号数量, right表示剩余可用的右括号数量
	backtrack = func(left, right int, track string) {
		if left == 0 && right == 0 { // 左括号的右括号都使用完毕，生成的肯定是合法的括号组合
			res = append(res, track)
			return
		}
		if right < left { // 目前生成的括号组合里，左括号的数量小于右括号的数量
			return
		}
		if left < 0 || right < 0 { 
			return
		}
		track += "("
		backtrack(left-1,right,track)
		track = track[:len(track)-1]

		track += ")"
		backtrack(left, right-1,track)
		track = track[:len(track)-1]
	}
	backtrack(n,n,"")
	return res
}
// @lc code=end

