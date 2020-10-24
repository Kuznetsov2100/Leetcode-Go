/*
 * @lc app=leetcode.cn id=1370 lang=golang
 *
 * [1370] 上升下降字符串
 */

// @lc code=start
func sortString(s string) string {
	// 题意解释：轮流按升序，降序添加字符到新的字符串中， 知道新字符串长度等于原字符串长度
	var res bytes.Buffer
	n := len(s)
	m := make([]int, 26)
	for i := range s {
		m[s[i]-'a']++
	}
	for res.Len() < n {
		for i := 0; i <= 25; i++ {
			if m[i] > 0 {
				res.WriteByte(byte(i+'a'))
				m[i]--
			}
		}
		for i := 25; i >= 0; i-- {
			if m[i] > 0 {
				res.WriteByte(byte(i+'a'))
				m[i]--
			}
		}	
	}
	return res.String()
}
// @lc code=end

