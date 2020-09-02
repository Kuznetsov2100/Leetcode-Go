/*
 * @lc app=leetcode.cn id=524 lang=golang
 *
 * [524] 通过删除字母匹配到字典里最长单词
 */

// @lc code=start
func findLongestWord(s string, d []string) string {
	sort.Slice(d, func(i, j int) bool {
		if m, n := len(d[i]),len(d[j]); m > n {
			return true
		} else if m == n {
			if d[i] < d[j] {
				return true
			}
		}
		return false
	})
	lens := len(s)
	for i := range d {
		length := len(d[i])
		if length > lens {
			continue
		}
		if isSubsequence(d[i], s, length, lens) {
			return d[i]
		}
	}
	return ""
}

func isSubsequence(s string, t string, lens, lent int) bool {
	indexS, indexT := 0, 0
	for indexS < lens && indexT < lent {
		if s[indexS] == t[indexT] {
			indexS++
		}
		indexT++
	}
	return indexS == lens
}
// @lc code=end

