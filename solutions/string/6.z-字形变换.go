/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 { // 行数小于2直接返回s
		return s
	}
	rows := make([]strings.Builder, numRows)
	i := 0
	flag := -1
	for j := range s {
		rows[i].WriteByte(s[j]) // 将字符s[j]填入对应行
		if i == 0 || i == numRows-1 { // 处于Z字形转折点时，方向反转
			flag = -flag
		}
		i += flag // 更新下一个字符对应行的索引
	}
	var res strings.Builder
	for _, row := range rows {
		res.WriteString(row.String()) // 拼接每一行的字符串得到结果字符串
	}
	return res.String()
}
// @lc code=end

