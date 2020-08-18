/*
 * @lc app=leetcode.cn id=1122 lang=golang
 *
 * [1122] 数组的相对排序
 */

// @lc code=start
func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := make(map[int]int)
	n := len(arr2)
	for i , v := range arr2 {
		m[v] = i
	}
	for _, v := range arr1 {
		if _, ok := m[v]; !ok {
			m[v] = n + v
		}
	}
	sort.Slice(arr1, func(i, j int) bool { return m[arr1[i]] < m[arr1[j]] })
	return arr1
}
// @lc code=end

