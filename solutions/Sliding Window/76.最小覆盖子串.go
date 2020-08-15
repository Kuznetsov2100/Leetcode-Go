/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
import "strings"
func minWindow(s string, t string) string {
	// 窗口[left,right)左闭右开
	left, right, match := 0, 0, 0
	start, minlen := 0, 1 << 63 - 1 // 记录最小子串在s中开始的索引及长度
	window := make(map[byte]int) // 记录窗口字符串中各字符出现的次数
	counter := make(map[byte]int) // 记录字符串t中各字符出现的次数
	for i := range t {
		counter[t[i]]++
	}
	for right < len(s) {
		char := s[right] // 可能要移入窗口的字符
		if _, ok := counter[char]; ok {
			window[char]++
			if window[char] == counter[char] {
				match++ // 如果s[right]在window和counter中出现的次数相同，则符合要求的字符数+1
			}
		}
		right++
		// 判断窗口左侧是否需要收缩
		for match == len(counter) {
			if right - left < minlen { // 如果找到一个更优解，更新最小字串在字符串s中开始的索引和长度
				start = left
				minlen = right - left
			}
			char := s[left] // 可能要移出窗口的字符
			if _, ok := counter[char]; ok {
				window[char]--
				if window[char] < counter[char] {
					match-- // 如果s[left]在window中出现的次数小于counter中的次数，则符合要求的字符数-1
				}
			}
			left++
		}
	}
	if minlen == (1 << 63 -1) { // s中不存在这样的子串
		return ""
	}
	return s[start:start+minlen] // 最小子串
}
// @lc code=end

