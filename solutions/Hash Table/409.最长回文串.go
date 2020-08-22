/*
 * @lc app=leetcode.cn id=409 lang=golang
 *
 * [409] 最长回文串
 */

// @lc code=start
func longestPalindrome(s string) int {
	// 对于一个回文串，最多只有一个字符出现奇数次，其余字符都出现偶数次
	var res int
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	for _, v := range m {
		if v > 1 {
			if v % 2 == 0 {  // 该字符出现偶数次，直接加总
				res += v
			} else { 
				res += v-1 // 该字符出现奇数次,该字符可利用的字符数为 v-1次
			}
		}
	}
	if res < len(s) &&  res % 2 == 0 { // 如果构造的最长回文串比s短，且回文串长度是偶数，那么回文串长度可以+1
		res++
	}
	return res
}
// @lc code=end

