/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 */

// @lc code=start
func countPrimes(n int) int {
	// Eratosthenes筛选法
	var res int
	prime := make([]bool, n)
	for i := range prime {
		prime[i] = true
	}
	for i := 2; i*i < n; i++ {
		if prime[i] {
			for j := i*i; j < n; j = j+i {
				prime[j] = false
			}
		}
	}
	for i := 2; i < n; i++ {
		if prime[i] {
			res++
		}
	}
	return res
}
// @lc code=end

