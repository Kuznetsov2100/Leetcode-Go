/*
 * @lc app=leetcode.cn id=383 lang=golang
 *
 * [383] 赎金信
 */

// @lc code=start
func canConstruct(ransomNote string, magazine string) bool {
	x, y := len(ransomNote), len(magazine)
	if y < x {
		return false
	}
	m := make(map[byte]int)
	for i := range magazine {
		m[magazine[i]]++
	}
	for i := range ransomNote {
		if _, ok := m[ransomNote[i]]; !ok {
			return false
		} else {
			m[ransomNote[i]]--
			if m[ransomNote[i]] < 0 {
				return false
			}
		}
	}
	return true
}
// @lc code=end

