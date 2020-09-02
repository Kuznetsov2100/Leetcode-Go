/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */

// @lc code=start
func maximumGap(nums []int) int {
	res := 0
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if diff := nums[i+1]-nums[i]; diff > res {
			res = diff
		}
	}
	return res
}
// @lc code=end

