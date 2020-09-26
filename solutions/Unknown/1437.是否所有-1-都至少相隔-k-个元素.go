/*
 * @lc app=leetcode.cn id=1437 lang=golang
 *
 * [1437] 是否所有 1 都至少相隔 k 个元素
 */

// @lc code=start
func kLengthApart(nums []int, k int) bool {
	pos1 := -1000000
	for i, num := range nums {
		if num == 1 {
			if i-pos1-1 < k {
				return false
			}
			pos1 = i
		}
	}
	return true
}
// @lc code=end

