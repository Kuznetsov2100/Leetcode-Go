/*
 * @lc app=leetcode.cn id=1232 lang=golang
 *
 * [1232] 缀点成线
 */

// @lc code=start
func checkStraightLine(coordinates [][]int) bool {
	dx := coordinates[0][0]-coordinates[1][0]
	dy := coordinates[0][1]-coordinates[1][1]	
	for i := 1; i < len(coordinates)-1; i++ {
		dxi := coordinates[i][0]-coordinates[i+1][0]
		dyi := coordinates[i][1]-coordinates[i+1][1]
		if dy * dxi != dx * dyi {
			return false
		}
	}
	return true
}
// @lc code=end

