/*
 * @lc app=leetcode.cn id=1185 lang=golang
 *
 * [1185] 一周中的第几天
 */

// @lc code=start
func dayOfTheWeek(day int, month int, year int) string {
	return time.Date(year, time.Month(month), day, 0,0,0,0, time.UTC).Weekday().String()
}
// @lc code=end

