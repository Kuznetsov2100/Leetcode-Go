/*
 * @lc app=leetcode.cn id=682 lang=golang
 *
 * [682] 棒球比赛
 */

// @lc code=start
func calPoints(ops []string) int {
	var res int
	var stack []int
	for i := range ops {
		if ops[i] == "C" {
			if len(stack) > 0 {
				res -= stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
		} else if ops[i] == "D" {
			num := stack[len(stack)-1]*2
			res += num
			stack = append(stack, num)
		} else if ops[i] == "+" {
			numA := stack[len(stack)-1]
			numB := stack[len(stack)-2]
			res += numA + numB
			stack = append(stack, numA+numB)
		} else {
			num, _ := strconv.Atoi(ops[i])
			res += num
			stack = append(stack, num)
		}
	}
	return res
}
// @lc code=end

