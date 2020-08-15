/*
 * @lc app=leetcode.cn id=76 lang=golang
 *
 * [76] 最小覆盖子串
 */

// @lc code=start
import "strings"
func minWindow(s string, t string) string {
	left, right, start, match := 0, 0, 0, 0
	minlen := 1 << 63 - 1 
	window := make(map[byte]int) // 记录窗口字符串中各字符出现的次数
	counter := make(map[byte]int) // 记录字符串t中各字符出现的次数
	for i := range t {
		counter[t[i]]++
	}
	for right < len(s) {
		char := s[right]
		if _, ok := counter[char]; ok { 
			window[char]++
			if window[char] == counter[char] {
				match++ // 如果s[right]在window和counter中出现的次数相同，则符合要求的字符数+1
			}
		}
		right++
		// 如果符合要求的字符数等于字符串t的长度，说明window中的字符串为一个可行解，现在要右移left缩小窗口优化可行解
		for match == len(counter) {
			if right - left < minlen { // 如果找到一个更优解，更新最小字串在字符串s中开始的索引和长度
				start = left
				minlen = right - left
			}
			char := s[left]
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

