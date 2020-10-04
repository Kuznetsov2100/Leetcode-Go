/*
 * @lc app=leetcode.cn id=599 lang=golang
 *
 * [599] 两个列表的最小索引总和
 */

// @lc code=start
func findRestaurant(list1 []string, list2 []string) []string {
	var res []string
	minIndexSum := math.MaxInt64
	m := make(map[string]int)
	for i, x := range list1 {
		m[x] = i
	}
	for i, y := range list2 {
		if index, ok := m[y]; ok {
			if index+i < minIndexSum {
				minIndexSum = index+i
				res = nil
				res = append(res, y)
			} else if index+i == minIndexSum {
				res = append(res, y)
			}
		}
	}
	return res
}
// @lc code=end

