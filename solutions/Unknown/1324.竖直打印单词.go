/*
 * @lc app=leetcode.cn id=1324 lang=golang
 *
 * [1324] 竖直打印单词
 */

// @lc code=start
func printVertically(s string) []string {
	var res []string
	var b bytes.Buffer
	strs := strings.Split(s, " ")
	n := len(strs)
	i := 0
	for {
		count := 0
		for j := 0; j < n; j++ {
			if i < len(strs[j]) {
				b.WriteByte(strs[j][i])
			} else {
				b.WriteByte(' ')
				count++
			}
		}
		if count == n {
			break
		}
		res = append(res, strings.TrimRight(b.String(), " "))
		b.Reset()
		i++
	}
	return res
}
// @lc code=end

