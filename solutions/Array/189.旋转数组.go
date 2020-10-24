/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 旋转数组
 */

// @lc code=start
func rotate(nums []int, k int)  {
	n := len(nums)
	k = k % n
	if k != 0 {
		res := make([]int, n)
		copy(res, nums)
		for i := 0; i < n; i++ {
			nums[i] = res[(n-k+i) % n]
		}
	}
}
// @lc code=end

