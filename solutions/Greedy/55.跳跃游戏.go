/*
 * @lc app=leetcode.cn id=55 lang=golang
 *
 * [55] 跳跃游戏
 */

// @lc code=start
func canJump(nums []int) bool {
	n , farthest := len(nums), 0
	for i := 0; i < n; i++ {
		if i <= farthest {
			farthest = max(farthest, i+nums[i])
			if farthest >= n-1 {
				return true
			}
		}
	}
	return false
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

