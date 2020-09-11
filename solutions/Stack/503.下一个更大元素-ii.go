/*
 * @lc app=leetcode.cn id=503 lang=golang
 *
 * [503] 下一个更大元素 II
 */

// @lc code=start
func nextGreaterElements(nums []int) []int {
	n := len(nums)
	var stack []int // 单调递减栈，栈中储存的是元素的是nums数组元素的索引
	res := make([]int, n)
	for i := range res {
		res[i] = -1
	}
	// nums数组中的每个元素要入栈两次，才能保证每个元素都有寻找下一个比它大的数的机会
	for i := 0; i < 2*n; i++ {
		j := i % n
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[j] {
			res[stack[len(stack)-1]] = nums[j]
			stack = stack[:len(stack)-1]		
		}
		stack = append(stack, j)	
	}
	return res
}
// @lc code=end

