/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 */

// @lc code=start
func isIsomorphic(s string, t string) bool {
	s_tmap := make(map[byte]byte)
	t_smap := make(map[byte]byte)
	for i := range s {
		s_t, ok1 := s_tmap[s[i]]
		t_s, ok2 := t_smap[t[i]]
		if (ok1 && s_t != t[i]) || (ok2 && t_s != s[i]) {
			return false
		} else {
			s_tmap[s[i]] = t[i]
			t_smap[t[i]] = s[i]
		}
	}	
	return true
}
// @lc code=end

