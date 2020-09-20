/*
 * @lc app=leetcode.cn id=395 lang=golang
 *
 * [395] 至少有K个重复字符的最长子串
 */

// @lc code=start
func longestSubstring(s string, k int) int {
	// 分治思想，首先统计各字符出现次数， 然后遍历s，遇到出现次数小于k的字符时，将s分为左右两部分子串递归计算
	// 右子串的起点为第一个出现次数 >= k的字符的索引
	// 假如字符串s长度 < k, 直接返回0
	// 假如字符串s中所有字符出现次数 >= k, 返回s的长度
	n := len(s)
	if n < k {
		return 0
	}
	m := make([]int, 26)
	for i := range s {
		m[s[i]-'a']++
	}
	for i := 0; i < n; i++ {
		if m[s[i]-'a'] < k {
			j := i
			for ; j < n; j++ {
				if m[s[j]-'a'] >= k {
					break
				}
			}
			return max(longestSubstring(s[:i], k), longestSubstring(s[j:], k))
		}
	}
	return n
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

