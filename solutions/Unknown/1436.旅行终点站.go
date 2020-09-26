/*
 * @lc app=leetcode.cn id=1436 lang=golang
 *
 * [1436] 旅行终点站
 */

// @lc code=start
func destCity(paths [][]string) string {
	var res string
	set := make(map[string]bool)
	for _, path := range paths {
		set[path[0]] = true
	}
	for _, path := range paths {
		if !set[path[1]] {
			res = path[1]
			break
		}
	}
	return res
}
// @lc code=end

