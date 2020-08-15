/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	// 窗口[left,right)左闭右开
	left, right, match := 0, 0, 0
	var res []int
	window := make(map[byte]int) // 记录窗口字符串中各字符出现的次数
	counter := make(map[byte]int) // 记录字符串t中各字符出现的次数
	for i := range p {
		counter[p[i]]++
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
			if right - left == len(p) {
				res = append(res,left)
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
	return res
}
// @lc code=end

