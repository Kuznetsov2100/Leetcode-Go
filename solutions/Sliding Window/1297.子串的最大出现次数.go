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
	window := make([]int, 26) // 记录窗口中子串各字符出现的频率
	windowLetters := 0
	for i := 0; i < minSize-1; i++ {
		window[s[i]-'a']++
	}

	for i := 0; i < 26; i++ {
		if window[i] > 0 {
			windowLetters++
		}
	}

	for i := 0; i <= n-minSize; i++ {
		left, right := i, i+minSize-1
		window[s[right]-'a']++
		if window[s[right]-'a'] == 1 {
			windowLetters++
		}
		if windowLetters <= maxLetters { // 窗口子串的不同字符个数小于等于maxLetters
			m[s[left:right+1]]++
			res = max(res, m[s[left:right+1]])
		}
		window[s[left]-'a']--
		if window[s[left]-'a'] == 0 { // 窗口最左侧字符出现次数降为0，则需从窗口中删除该字符
			windowLetters--
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

