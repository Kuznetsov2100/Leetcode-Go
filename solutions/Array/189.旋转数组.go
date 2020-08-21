/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	n := len(nums)
	if k == n {
		return
	} else if k > n {
		k = k % n
	}
	copy(nums,  append(nums[n-k:],nums[:n-k]...))
}
// @lc code=end

