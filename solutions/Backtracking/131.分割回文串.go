/*
 * @lc app=leetcode.cn id=131 lang=golang
 *
 * [131] 分割回文串
 */

// @lc code=start
func partition(s string) [][]string {
	var res [][]string
	length := len(s)
	if length == 0 {
		return res
	}
	isPalindrome := func(i int, j int) bool {
		for i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
		return true
	}
	var backtrack func(pos int, substrlen int, track []string)
	backtrack = func(pos int, substrlen int, track []string) {
		if substrlen == 0 {
			tmp := make([]string, len(track))
			copy(tmp,track)
			res = append(res, tmp)
			return
		}
		for i := 1; i <= substrlen; i++ {
			if i > 1 && !isPalindrome(pos,pos+i-1) {
				continue
			}
			track = append(track, s[pos:pos+i])
			backtrack(pos+i,substrlen-i,track)
			track = track[:len(track)-1]
		}
	}
	backtrack(0,length,[]string{})
	return res
}
// @lc code=end

