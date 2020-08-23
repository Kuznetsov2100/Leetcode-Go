/*
 * @lc app=leetcode.cn id=38 lang=golang
 *
 * [38] 外观数列
 */

// @lc code=start
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	str := countAndSay(n-1)
	var res strings.Builder
	start := 0
	for i := 1; i < len(str)+1; i++ {
		if i == len(str) {
			res.WriteString(strconv.Itoa(i-start))
			res.WriteByte(str[start])
		} else {
			if str[i] != str[start] {
				res.WriteString(strconv.Itoa(i-start))
				res.WriteByte(str[start])
				start = i
			}
		}
	}
	return res.String()
}
// @lc code=end

