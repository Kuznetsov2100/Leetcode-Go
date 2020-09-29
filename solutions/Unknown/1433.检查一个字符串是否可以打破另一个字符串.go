/*
 * @lc app=leetcode.cn id=1433 lang=golang
 *
 * [1433] 检查一个字符串是否可以打破另一个字符串
 */

// @lc code=start
func checkIfCanBreak(s1 string, s2 string) bool {
	s1b := []byte(s1)
	s2b := []byte(s2)
	sort.Slice(s1b, func(i, j int) bool {return s1b[i] < s1b[j]})
	sort.Slice(s2b, func(i, j int) bool {return s2b[i] < s2b[j]})
	greatercount, smallercount, n := 0, 0, len(s1)
	for i := 0; i < n; i++ {
		if s1b[i] <= s2b[i] {
			smallercount++
		}
		if s1b[i] >= s2b[i] {
			greatercount++
		}
	}
	if smallercount == n || greatercount == n {
		return true
	}
	return false
}
// @lc code=end

