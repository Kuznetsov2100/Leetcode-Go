/*
 * @lc app=leetcode.cn id=821 lang=golang
 *
 * [821] 字符的最短距离
 */

// @lc code=start
func shortestToChar(S string, C byte) []int {
	// 对于S中每个字符c, 维护两个变量：
	// leftIndex: 字符c左边最近的字符C的索引
	// rightIndex: 字符c右边最近的字符C的索引
	res := make([]int, len(S))
	leftIndex := -100000
	rightIndex := strings.IndexByte(S, C)
	for i := range S {
		if S[i] == C {
			res[i] = 0
			leftIndex = i
			if index := strings.IndexByte(S[i+1:], C); index == -1 {
				rightIndex = 100000
			} else {
				rightIndex = i+index+1
			}
		} else {
			res[i] = min(rightIndex-i, i-leftIndex)
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

