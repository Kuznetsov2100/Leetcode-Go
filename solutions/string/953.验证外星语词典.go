/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	alpha := make([]int, 26)
	for i := range order {
		alpha[order[i]-'a'] = i
	}
	isSmaller := func(a, b string) bool {
		m, n := len(a), len(b)
		length := min(m,n)
		for i := 0; i < length; i++ {
			if alpha[a[i]-'a'] < alpha[b[i]-'a'] {
				return true
			} else if alpha[a[i]-'a'] > alpha[b[i]-'a'] {
				return false
			}
		}
		return m < n
	}
	return sort.SliceIsSorted(words, func(i, j int) bool { return isSmaller(words[i], words[j]) })
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

