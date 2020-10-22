/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	m, n := len(str1), len(str2)
	commonDivisor := gcd(m, n)
	res := str1[:commonDivisor]
	if strings.Repeat(res, m/commonDivisor) != str1 || strings.Repeat(res, n/commonDivisor) != str2 {
		return ""
	}
	return res
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}
// @lc code=end

