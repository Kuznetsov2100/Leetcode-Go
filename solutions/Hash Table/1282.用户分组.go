/*
 * @lc app=leetcode.cn id=1282 lang=golang
 *
 * [1282] 用户分组
 */

// @lc code=start
func groupThePeople(groupSizes []int) [][]int {
	var res [][]int
	m := make(map[int][]int)
	for i, size := range groupSizes {
		m[size] = append(m[size], i)
	}
	for size, ids := range m {
		for i := 0; i+size <= len(ids); i += size {
			res = append(res, ids[i:i+size])
		}
	}
	return res
}
// @lc code=end

