/*
 * @lc app=leetcode.cn id=50 lang=golang
 *
 * [50] Pow(x, n)
 */

// @lc code=start
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

// 递归版本
// func myPow(x float64, n int) float64 {
// 	if n == 0 {
// 		return 1.0
// 	}
// 	if n < 0 {
// 		x = 1/x
// 		n = -n
// 	}
// 	res := myPow(x, n >> 1)
// 	if n & 1 == 1 {
// 		return res * res * x
// 	} else {
// 		return res * res
// 	}
// }
// @lc code=end

