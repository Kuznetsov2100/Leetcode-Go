/*
 * @lc app=leetcode.cn id=290 lang=golang
 *
 * [290] 单词规律
 */

// @lc code=start
import "strings"
func wordPattern(pattern string, str string) bool {
	arr := strings.Split(str, " ")
	if len(arr) != len(pattern) {
		return false
	}
	ptosmap := make(map[byte]string)
	stopmap := make(map[string]byte)
	for i := range pattern {
		ptos, ok1 := ptosmap[pattern[i]]
		stop, ok2 := stopmap[arr[i]]
		if (ok1 && ptos != arr[i]) || (ok2 && stop != pattern[i]) {
			return false
		} else {
			ptosmap[pattern[i]] = arr[i]
			stopmap[arr[i]] = pattern[i]
		}
	}	
	return true
}
// @lc code=end

