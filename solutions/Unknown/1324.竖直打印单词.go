/*
 * @lc app=leetcode.cn id=1324 lang=golang
 *
 * [1324] 竖直打印单词
 */

// @lc code=start
func printVertically(s string) []string {
	var res []string
	var b bytes.Buffer
	var maxlen int
	strs := strings.Split(s, " ")
	for i := range strs {
		if x := len(strs[i]); x > maxlen {
			maxlen = x
		}
	}
	
	for i := 0; i < maxlen; i++ {
		for j := 0; j < len(strs); j++ {
			if i < len(strs[j]) {
				b.WriteByte(strs[j][i])
			} else {
				b.WriteByte(' ')
			}
		}
		res = append(res, strings.TrimRight(b.String(), " "))
		b.Reset()
	}
	return res
}
// @lc code=end

