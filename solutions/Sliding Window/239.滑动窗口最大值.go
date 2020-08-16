/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return nil
	}
	res := make([]int,len(nums)-k+1)
	deque := make([]int,0,k)
	for i := range nums {
		for i > 0 && len(deque) > 0 && nums[i] > nums[deque[len(deque)-1]] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i-k >= 0  && deque[0] <= i-k {
			deque = deque[1:]
		}
		if i-k + 1 >= 0 {
			res[i-k+1] = nums[deque[0]]
		}
	}
	return res
}
// @lc code=end

