/*
 * @lc app=leetcode.cn id=372 lang=golang
 *
 * [372] 超级次方
 */

// @lc code=start
func superPow(a int, b []int) int {
	n := len(b)
	if n == 0 {
		return 1
	}
	part1 := myPow(a, b[n-1])
	part2 := myPow(superPow(a, b[:n-1]), 10)
	return (part1 * part2) % 1337
}

func myPow(x int, n int) int {
	if n == 0 {
		return 1
	}
	x = x % 1337
	res := myPow(x, n >> 1)
	if n & 1 == 1 {
		return (res * res * x) % 1337
	} else {
		return (res * res) % 1337
	}
}
// @lc code=end

