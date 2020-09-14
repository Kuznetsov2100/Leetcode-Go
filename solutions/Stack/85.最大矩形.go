/*
 * @lc app=leetcode.cn id=85 lang=golang
 *
 * [85] 最大矩形
 */

// @lc code=start
func maximalRectangle(matrix [][]byte) int {
	// 将矩阵的每一行看作是柱形图，某个柱形的高度为某列截止到当前行的连续'1'的个数，
	// 所有柱形的宽度为1，调用84题的解法,更新最大面积
	var res int
	if len(matrix) == 0 {
		return res
	}
	row, col := len(matrix), len(matrix[0])
	heights := make([]int, col)
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		ans := largestRectangleArea(heights)
		if ans > res {
			res = ans
		}
	}
	return res
}

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

