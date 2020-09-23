/*
 * @lc app=leetcode.cn id=492 lang=golang
 *
 * [492] 构造矩形
 */

// @lc code=start
func constructRectangle(area int) []int {
	res := []int{area, area}
	for W := mySqrt(area); W >= 1; W-- {
		if area % W == 0 {
			res = []int{area/W, W}
			break
		}
	}
	return res
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

