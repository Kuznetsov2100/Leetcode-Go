/*
 * @lc app=leetcode.cn id=84 lang=golang
 *
 * [84] 柱状图中最大的矩形
 */

// @lc code=start
func largestRectangleArea(heights []int) int {
	// 维持一柱形高度递增的单调递增栈
	var res int
	var stack []int // 储存的是柱形高度在数组中的索引
	// 在柱体数组的头和尾加了两个高度为0的柱形
	// 最左边的高度为0柱形由于它一定比输入数组里任何一个元素小，它肯定不会出栈，因此栈一定不会为空；
    // 最右边的高度为0柱形一定比输入数组里任何一个元素小，它会让所有输入数组里的元素出栈（最左边高度为0柱形除外）
	heights = append(append([]int{0}, heights...), 0)
	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			n := len(stack)
			hindex := stack[n-1]
			left := stack[n-2] + 1
			stack = stack[:n-1]
			right := i
			if ans := heights[hindex] * (right-left); ans > res {
				res = ans
			}
		}
		stack = append(stack, i)
	}
	return res
}
// @lc code=end

