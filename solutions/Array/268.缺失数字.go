/*
 * @lc app=leetcode.cn id=268 lang=golang
 *
 * [268] 缺失数字
 */

// @lc code=start
func missingNumber(nums []int) int {
	// 求等差数列0,1,2,...,n的和， 减去nums数组的和，即为缺失的数字
	n := len(nums)
	nplus1sum := (0+n)*(n+1)/2
	nsum := 0
	for i := range nums {
		nsum += nums[i]
	}
	return nplus1sum-nsum
}
// @lc code=end

