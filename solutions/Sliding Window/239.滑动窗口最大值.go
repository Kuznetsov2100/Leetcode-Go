/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */

// @lc code=start
func maxSlidingWindow(nums []int, k int) []int {
	// 1. deque最左侧元素保存着当前窗口最大值的下标。
	// 2. 当我们遇到新的数时，将新的数和deque的末尾比较，如果末尾比新数小，则把末尾扔掉，
	//    直到该队列的末尾比新数大或者队列为空的时候才停止。这样保证deque是一个单调数列。
	// 3. deque的所有值要在窗口范围内
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
		if index := i - k + 1; index >= 0 { // deque初始化完毕
			res[index] = nums[deque[0]] // 将当前窗口的最大值加入结果数组
			if deque[0] <= index { // 保证deque中的索引在当前窗口范围内
				deque = deque[1:]
			}
		}
	}
	return res
}
// @lc code=end

