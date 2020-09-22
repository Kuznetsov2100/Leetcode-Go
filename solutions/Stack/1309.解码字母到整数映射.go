/*
 * @lc app=leetcode.cn id=1309 lang=golang
 *
 * [1309] 解码字母到整数映射
 */

// @lc code=start
func freqAlphabets(s string) string {
	var stack []byte
	for i := range s {
		if s[i] == '#' {
			number, _ := strconv.Atoi(string(stack[len(stack)-2:]))
			stack = stack[:len(stack)-2]
			stack = append(stack, byte(number-10)+'j')
		} else {
			stack = append(stack, s[i])
		}
	}
	for i := range stack {
		if '0' <= stack[i] && stack[i] <= '9' {
			stack[i] = stack[i] - '1' + 'a'
		}
	}
	return string(stack)
}
// @lc code=end

