/*
 * @lc app=leetcode.cn id=423 lang=golang
 *
 * [423] 从英文中重建数字
 */

// @lc code=start
func originalDigits(s string) string {
	var res strings.Builder
	num := make([]int, 10)
	m := make(map[byte]int)
	for i := range s {
		m[s[i]]++
	}
	num[0] = m['z']
	num[2] = m['w']
	num[4] = m['u']
	num[6] = m['x']
	num[8] = m['g']
	num[3] = m['h'] - num[8]
	num[5] = m['f'] - num[4]
	num[7] = m['s'] - num[6]
	num[9] = m['i'] - num[5] - num[6] - num[8]
	num[1] = m['o'] - num[0] - num[2] - num[4]
	for i := range num {
		for j := 0; j < num[i]; j++ {
			res.WriteString(strconv.Itoa(i))
		}
	}
	return res.String()
}
// @lc code=end

