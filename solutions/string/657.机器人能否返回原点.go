/*
 * @lc app=leetcode.cn id=657 lang=golang
 *
 * [657] 机器人能否返回原点
 */

// @lc code=start
func judgeCircle(moves string) bool {
	x, y := 0, 0
	for i := range moves {
		if moves[i] == 'U' {
			y++
		} else if moves[i] == 'L' {
			x--
		} else if moves[i] == 'R' {
			x++
		} else if moves[i] == 'D' {
			y--
		}
	}
	return x == 0 && y == 0
}
// @lc code=end

