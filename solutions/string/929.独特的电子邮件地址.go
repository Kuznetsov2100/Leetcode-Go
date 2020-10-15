/*
 * @lc app=leetcode.cn id=929 lang=golang
 *
 * [929] 独特的电子邮件地址
 */

// @lc code=start
func numUniqueEmails(emails []string) int {
	var res int
	var b bytes.Buffer
	m := make(map[string]bool)
	for _, email := range emails {	
		var index int
		indexPlus := strings.Index(email, "+")
		if indexPlus == -1 {
			indexPlus = 100 // 如果email不含"+", 则将索引设为最大值
		}
		for i := range email {
			if email[i] == '@' {
				index = i
				break
			}
			if email[i] == '.' || i >= indexPlus {
				continue
			}
			b.WriteByte(email[i])
		}
		b.WriteString(email[index:])
		email = b.String()
		b.Reset()
		if !m[email] {
			m[email] = true
			res++
		}
	}
	return res	
}
// @lc code=end

