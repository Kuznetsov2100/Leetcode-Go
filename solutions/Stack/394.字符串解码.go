/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */

// @lc code=start
func decodeString(s string) string {
	var res,RepeatNumber bytes.Buffer
	var alphaStack []string // 储存'['和字母的栈
	var numStack []int // 储存重复次数的栈
	n := len(s)
	for i := 0; i < n; i++ {
		if isAlpha(s[i]) {
			if len(alphaStack) == 0 {
				res.WriteByte(s[i]) // 字母栈为空，s[i]直接写入结果
			} else {
				alphaStack = append(alphaStack, string(s[i]))
			}
		} else if isDigit(s[i]){
			for i < n && isDigit(s[i]) {
				RepeatNumber.WriteByte(s[i])
				i++
			}
			num, _ := strconv.Atoi(RepeatNumber.String())
			numStack = append(numStack, num)
			RepeatNumber.Reset()
			i-- // 需要后退一步
		} else if s[i] == '[' {
			alphaStack = append(alphaStack, string(s[i]))
		} else if s[i] == ']' {
			var alpha string // alpha储存的是alphaStack出栈的字符
			for len(alphaStack) > 0 && alphaStack[len(alphaStack)-1] != "[" {
				alpha = alphaStack[len(alphaStack)-1] + alpha
				alphaStack = alphaStack[:len(alphaStack)-1]
			}
			alphaStack = alphaStack[:len(alphaStack)-1] // pop "["

			str := strings.Repeat(alpha, numStack[len(numStack)-1]) // 计算中间结果
			numStack = numStack[:len(numStack)-1]
			if len(alphaStack) > 0 {
				alphaStack = append(alphaStack, str) // 字母栈非空，需要将中间结果入栈
			} else {
				res.WriteString(str) // 字母栈为空， 直接写入最终结果
			}
		}
	}
	return res.String()
}

// 判断字符是否为字母
func isAlpha(x byte) bool {
	if ('A' <= x && x <= 'Z') || ('a' <= x && x <= 'z') {
		return true
	}
	return false
}

// 判断字符是否为数字
func isDigit(x byte) bool {
	if '0' <= x && x <= '9' {
		return true
	}
	return false
}
// @lc code=end

