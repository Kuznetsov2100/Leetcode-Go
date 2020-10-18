/*
 * @lc app=leetcode.cn id=594 lang=golang
 *
 * [594] 最长和谐子序列
 */

// @lc code=start
func findLHS(nums []int) int {
	m := make(map[int]int)
	for i := range nums {
		m[nums[i]]++
	}
	var res int
	for k, v := range m {
		if count, ok := m[k+1]; ok {
			res = max(res, v+count)
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

