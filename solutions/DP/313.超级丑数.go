/*
 * @lc app=leetcode.cn id=313 lang=golang
 *
 * [313] 超级丑数
 */

// @lc code=start
// 动态规划
func nthSuperUglyNumber(n int, primes []int) int {
	p := make([]int, len(primes)) // 多指针
	nums := make([]int, n)
	for i := range nums {
		nums[i] = 1 << 31-1
	}
	nums[0] = 1 // base case
	for i := 1; i < n; i++ {
		for j, prime := range primes {
			nums[i] = min(nums[i], prime*nums[p[j]])
		}
		for k, prime := range primes {
			if nums[i] == nums[p[k]]*prime {
				p[k]++
			}
		}
	}
	return nums[n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

