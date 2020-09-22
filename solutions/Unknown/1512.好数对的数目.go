/*
 * @lc app=leetcode.cn id=1512 lang=golang
 *
 * [1512] 好数对的数目
 */

// @lc code=start
func numIdenticalPairs(nums []int) int {
	var res int
	m := make([]int, 101)
	for _, num := range nums {
		m[num]++
	}
	for _, count := range m {
		if count > 1 {
			res += count*(count-1)/2
		}
	}
	return res
}
// @lc code=end

