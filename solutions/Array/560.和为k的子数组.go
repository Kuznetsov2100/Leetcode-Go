/*
 * @lc app=leetcode.cn id=560 lang=golang
 *
 * [560] 和为K的子数组
 */

// @lc code=start
func subarraySum(nums []int, k int) int {
	var res, sum int
	m := make(map[int]int)
	m[0] = 1
	for i := range nums {
		sum += nums[i]
		if count, ok := m[sum-k]; ok {
			res += count
		}
		m[sum]++
	}
	return res
}
// @lc code=end

