/*
 * @lc app=leetcode.cn id=747 lang=golang
 *
 * [747] 至少是其他数字两倍的最大数
 */

// @lc code=start
func dominantIndex(nums []int) int {
	var res int
	first, second := -1, -1
	for i, v := range nums {
		if v > first {
			first, second = v, first
			res = i
		} else if v > second {
			second = v
		}
	}
	if second  * 2 <= first {
		return res
	}
	return -1
}
// @lc code=end

