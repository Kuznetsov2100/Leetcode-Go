/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel表列序号
 */

// @lc code=start
func titleToNumber(s string) int {
	var res int
	for i, j := 0, len(s)-1; i < len(s) && j >= 0; i,j = i+1,j-1 {
		res += int(myPow(26, j))*int(s[i]-64)
	}
	return res	 
}

func myPow(x float64, n int) float64 {
	// 快速幂算法
	if n == 0 {
		return 1.0
	}
	if n < 0 {
		n = -n
		x = 1/x
	}
	res := 1.0
	for n > 0 {
		if n & 1 == 1 { // 如果n为奇数,需要多乘以一个x
			res *= x
		}
		x *= x
		n >>= 1 // n/2 
	}
	return res
}
// @lc code=end

