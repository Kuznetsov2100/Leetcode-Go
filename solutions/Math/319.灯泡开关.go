/*
 * @lc app=leetcode.cn id=319 lang=golang
 *
 * [319] 灯泡开关
 */

// @lc code=start
func bulbSwitch(n int) int {
	return mySqrt(n)
}

func mySqrt(x int) int {
	left, right := 0, x
	res := -1
	for left <= right {
		mid := left + (right - left)/2
		if mid * mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return res
}
// @lc code=end

