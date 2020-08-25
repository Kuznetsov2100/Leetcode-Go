/*
 * @lc app=leetcode.cn id=11 lang=golang
 *
 * [11] 盛最多水的容器
 */

// @lc code=start
func maxArea(height []int) int {
	var res int
	left, right := 0, len(height)-1
	for left < right {
		if height[left] < height[right] {
			res = max(res, height[left]*(right-left))
			left++ // 尽可能提升短板，才有机会获得更大的容量
		} else {
			res = max(res, height[right]*(right-left))
			right--
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

