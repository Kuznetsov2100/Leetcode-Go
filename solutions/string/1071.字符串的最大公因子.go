/*
 * @lc app=leetcode.cn id=1071 lang=golang
 *
 * [1071] 字符串的最大公因子
 */

// @lc code=start
func gcdOfStrings(str1 string, str2 string) string {
	if str1 + str2 != str2 + str1 {
		return ""
	}
	return str1[:gcd(len(str1), len(str2))]
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}
// @lc code=end

