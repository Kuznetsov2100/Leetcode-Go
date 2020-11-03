/*
 * @lc app=leetcode.cn id=1297 lang=golang
 *
 * [1297] 子串的最大出现次数
 */

// @lc code=start
func maxFreq(s string, maxLetters int, minSize int, maxSize int) int {
	// 只需考虑长度为minSize的符合条件的子串
	var res int
	n := len(s)
	m := make(map[string]int)
	window := make(map[byte]int) // 记录窗口中子串各字符出现的频率
	for i := 0; i < minSize-1; i++ {
		window[s[i]]++
	}

	for i := 0; i <= n-minSize; i++ {
		candidate := s[i:i+minSize] // 当前候选子串
		window[s[i+minSize-1]]++
		if len(window) <= maxLetters { // 窗口子串的不同字符个数小于等于maxLetters
			m[candidate]++
			res = max(res, m[candidate])
		}
		window[s[i]]--
		if window[s[i]] == 0 { // 窗口最左侧字符出现次数降为0，则需从窗口中删除该字符
			delete(window, s[i])
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

