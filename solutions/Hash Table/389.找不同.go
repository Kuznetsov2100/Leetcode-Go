/*
 * @lc app=leetcode.cn id=389 lang=golang
 *
 * [389] 找不同
 */

// @lc code=start
func findTheDifference(s string, t string) byte {
	var res byte
	smap := [26]int{}
	for i := range s {
		smap[s[i]-'a']++
	}
	for i := range t {
		smap[t[i]-'a']--
	}
	for i := range smap {
		if smap[i] < 0 {
			res = byte(i+97)
			break
		}
	}
	return res
}
// @lc code=end

