/*
 * @lc app=leetcode.cn id=1296 lang=golang
 *
 * [1296] 划分数组为连续数字的集合
 */

// @lc code=start
func isPossibleDivide(nums []int, k int) bool {
	if len(nums) % k != 0 {
		return false
	}
	sort.Ints(nums)
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	for i := range nums {
		if m[nums[i]] > 0 {
			for j := 0; j < k; j++ {
				m[nums[i]+j]--
				if m[nums[i]+j] < 0 {
					return false
				}
			}
		}
	}
	return true	
}
// @lc code=end

