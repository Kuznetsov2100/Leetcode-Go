/*
 * @lc app=leetcode.cn id=844 lang=golang
 *
 * [844] 比较含退格的字符串
 */

// @lc code=start
func backspaceCompare(S string, T string) bool {
	var stackS, stackT []byte
	for i := range S {
		if S[i] == '#' {
			if len(stackS) > 0 {
				stackS = stackS[:len(stackS)-1]
			}
		} else {
			stackS = append(stackS, S[i])
		}
	}
	for i := range T {
		if T[i] == '#' {
			if len(stackT) > 0 {
				stackT = stackT[:len(stackT)-1]
			}
		} else {
			stackT = append(stackT, T[i])
		}
	}
	return string(stackS) == string(stackT)
}
// @lc code=end

