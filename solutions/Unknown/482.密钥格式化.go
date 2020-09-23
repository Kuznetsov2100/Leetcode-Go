/*
 * @lc app=leetcode.cn id=482 lang=golang
 *
 * [482] 密钥格式化
 */

// @lc code=start
func licenseKeyFormatting(S string, K int) string {
	var res bytes.Buffer
	S = strings.ReplaceAll(S, "-", "")
	n, part1Len, cnt := len(S), 1, 0
	
	for part1Len < n {
		if (n-part1Len) % K == 0 {
			break
		}
		part1Len++
	}

	for i := range S {
		ch := S[i]
		if 'a' <= ch && ch <= 'z' {
			ch -= 32
		}
		if i >= part1Len {
			if cnt % K == 0 {
				res.WriteByte('-')
			}
			cnt++
		}	
		res.WriteByte(ch)
	}
	return res.String()
}
// @lc code=end

