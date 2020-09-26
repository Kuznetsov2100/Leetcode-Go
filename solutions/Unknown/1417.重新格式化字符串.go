/*
 * @lc app=leetcode.cn id=1417 lang=golang
 *
 * [1417] 重新格式化字符串
 */

// @lc code=start
func reformat(s string) string {
	var alpha, num []byte
	for i := range s {
		if 'a' <= s[i] && s[i] <= 'z' {
			alpha = append(alpha, s[i])
		} else {
			num = append(num, s[i])
		}
	}

	m, n := len(alpha), len(num)
	if m-n > 1 || n-m > 1 {
		return ""
	}
	
	length := 0
	if m < n {
		length = m
	} else {
		length = n
	}
	var res bytes.Buffer
	for i := 0; i < length; i++ {
		if m >= n {
			res.WriteByte(alpha[i])
			res.WriteByte(num[i])
		} else {
			res.WriteByte(num[i])
			res.WriteByte(alpha[i])
		}
	}
	if m > n {
		res.WriteByte(alpha[m-1])
	} else if m < n {
		res.WriteByte(num[n-1])
	}
	return res.String()
}
// @lc code=end

