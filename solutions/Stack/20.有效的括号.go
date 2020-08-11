/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

// @lc code=start
func isValid(s string) bool {
	validpair := map[byte]byte{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	var stackA,stackB []byte
	for i := range s {
		stackA = append(stackA, s[i])
	}
	for len(stackA) > 0 {
		left := stackA[len(stackA)-1]
		stackA = stackA[:len(stackA)-1]
		if left == ')'  || left == '}' || left == ']' {
			stackB = append(stackB, left)
		} else {
			if len(stackB) == 0 {
				return false
			}
			right := stackB[len(stackB)-1]
			stackB = stackB[:len(stackB)-1]
			if validpair[left] != right {
				return false
			}
		}
	}
	if len(stackB) > 0 {
		return false
	}
	return true
}
// @lc code=end

