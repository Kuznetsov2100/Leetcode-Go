/*
 * @lc app=leetcode.cn id=1346 lang=golang
 *
 * [1346] 检查整数及其两倍数是否存在
 */

// @lc code=start
func checkIfExist(arr []int) bool {
	m := make(map[int]int)
	for i := range arr {
		m[arr[i]]++
	}
	for _, num := range arr {
		if num != 0 && m[2*num] >= 1 {
			return true
		}
		if num == 0 && m[0] >= 2 {
			return true
		}
	}
	return false
}
// @lc code=end

