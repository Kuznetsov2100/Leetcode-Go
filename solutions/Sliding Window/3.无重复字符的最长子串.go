/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	// 窗口[left,right)左闭右开
	left, right,res := 0, 0, 0
	window := make(map[byte]int) // 记录窗口字符串中各字符出现的次数
	for right < len(s) {
		char := s[right]
		window[char]++
		right++
		// 如果window出现重复字符，开始缩小窗口左侧
		for window[char] > 1 {
			window[s[left]]--
			left++
		}
		res = max(res,right-left)	
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

