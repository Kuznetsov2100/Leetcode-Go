/*
 * @lc app=leetcode.cn id=202 lang=golang
 *
 * [202] 快乐数
 */

// @lc code=start
func isHappy(n int) bool {
	slow, fast := n, step(n)
	for fast != 1 {
		slow = step(slow)
		fast = step(step(fast))
		if slow == fast {
			return false
		}
	}
	return true
}

func step(n int) int {
	sum := 0
	for n > 0 {
		sum += (n%10) * (n%10)
		n /= 10
	}
	return sum
}
// @lc code=end

