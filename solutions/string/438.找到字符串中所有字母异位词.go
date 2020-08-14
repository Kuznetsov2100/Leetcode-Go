/*
 * @lc app=leetcode.cn id=438 lang=golang
 *
 * [438] 找到字符串中所有字母异位词
 */

// @lc code=start
func findAnagrams(s string, p string) []int {
	lens,lenp := len(s),len(p)
	if lens < lenp {
		return nil
	}
	var res []int
	for i := 0; i < lens-lenp+1; i++ {
		if isAnagram(s[i:lenp+i],p) {
			res = append(res, i)
		}
	}
	return res
}

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make([]int,26)
	for i := 0; i < len(s); i++ {
		count[s[i]-97]++
		count[t[i]-97]--
	}
	for _, v := range count {
		if v != 0 {
			return false
		}
	}
	return true
}
// @lc code=end

