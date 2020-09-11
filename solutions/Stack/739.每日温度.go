/*
 * @lc app=leetcode.cn id=739 lang=golang
 *
 * [739] 每日温度
 */

// @lc code=start
func dailyTemperatures(T []int) []int {
	n := len(T)
	res := make([]int, n)
	var stack []int // 单调递减栈， 栈中储存的是数组T中元素的索引
	for i, temperature := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < temperature {
			index := stack[len(stack)-1]
			res[index] = i-index
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return res
}
// @lc code=end

