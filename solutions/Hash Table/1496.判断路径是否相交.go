/*
 * @lc app=leetcode.cn id=1496 lang=golang
 *
 * [1496] 判断路径是否相交
 */

// @lc code=start
func isPathCrossing(path string) bool {
	type point [2]int
	x, y := 0, 0
	m := map[point]bool{ point{x, y} : true }
	for i := range path {
		if path[i] == 'N' {
			y++
		} else if path[i] == 'S' {
			y--
		} else if path[i] == 'E' {
			x++
		} else if path[i] == 'W' {
			x--
		}
		if m[point{x, y}] {
			return true
		}
		m[point{x, y}] = true
	}
	return false
}
// @lc code=end

