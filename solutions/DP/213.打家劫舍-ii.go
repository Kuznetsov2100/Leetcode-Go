/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	// 有三种情况：
	// case1: 偷第一家，不偷最后一家 helper(nums[:n-1], n-1)
	// case2: 不偷第一家， 偷最后一家 helper(nums[1:], n-1)
	// case3: 既不偷第一家， 也不偷最后一家
	// 很明显case3的结果最小，只需比较case1和case2谁更大
	return max(helper(nums[:n-1], n-1), helper(nums[1:], n-1))
}

func helper(nums []int, length int) int {
	dp := make([]int, length+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i <= length; i++ {
		dp[i] = max( dp[i-1], dp[i-2] + nums[i-1])
	}
	return dp[length]
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

