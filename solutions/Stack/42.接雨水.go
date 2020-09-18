/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {
	var res int
	var stack []int // 单调递减栈，栈中保存的是height数组中的元素的索引
	for i := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] < height[i] {
			bottom := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			width := i-stack[len(stack)-1]-1 // 当前可计算的雨水面积的宽度
			height := min(height[i], height[stack[len(stack)-1]])-bottom // 高度
			res += width*height
		}
		stack = append(stack, i)
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

