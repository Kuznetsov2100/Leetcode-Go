/*
 * @lc app=leetcode.cn id=1450 lang=golang
 *
 * [1450] 在既定时间做作业的学生人数
 */

// @lc code=start
func busyStudent(startTime []int, endTime []int, queryTime int) int {
	var res int
	for i := range startTime {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			res++
		}
	}
	return res
}
// @lc code=end

