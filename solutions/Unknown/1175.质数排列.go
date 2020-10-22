/*
 * @lc app=leetcode.cn id=1175 lang=golang
 *
 * [1175] 质数排列
 */

// @lc code=start
func numPrimeArrangements(n int) int {
	primeCount := countPrimes(n+1)
	return (factorial(primeCount) * factorial(n-primeCount)) % 1000000007
}

func factorial(n int) int {
	res := 1
	for i :=2; i <= n; i++ {
		res = (res * i) % 1000000007
	}
	return res
}

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

