/*
 * @lc app=leetcode.cn id=1002 lang=golang
 *
 * [1002] 查找常用字符
 */

// @lc code=start
func commonChars(A []string) []string {
	var res []string
	n := len(A)
	matrix := make([][]int, 26)
	for i := range matrix {
		matrix[i] = make([]int, n)
	}
	for i, str := range A {
		for j := range str {
			matrix[str[j]-'a'][i]++
		}
	}
	for i := 0; i < 26; i++ {
		count := 101
		for j := 0; j < n; j++ {
			if matrix[i][j] < count {
				count = matrix[i][j]
			}
		}
		for count > 0 {
			res = append(res, string(i+97))
			count--
		}
	}
	return res
}
// @lc code=end

