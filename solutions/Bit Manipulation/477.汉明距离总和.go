/*
 * @lc app=leetcode.cn id=477 lang=golang
 *
 * [477] 汉明距离总和
 */

// @lc code=start
func totalHammingDistance(nums []int) int {
	var res int
	n := len(nums)
	for i := 0; i < 32; i++ {
		pos := 0
		for j := range nums {
			if (nums[j] >> i) & 1 == 1 {
				pos++
			}
		}
		res += pos * (n-pos)
	}
	return res
}
// @lc code=end

