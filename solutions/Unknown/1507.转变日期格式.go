/*
 * @lc app=leetcode.cn id=1507 lang=golang
 *
 * [1507] 转变日期格式
 */

// @lc code=start
func reformatDate(date string) string {
	m := map[string]string{
		"Jan": "01",
		"Feb": "02",
		"Mar": "03",
		"Apr": "04",
		"May": "05",
		"Jun": "06",
		"Jul": "07",
		"Aug": "08",
		"Sep": "09",
		"Oct": "10",
		"Nov": "11",
		"Dec": "12",
	}
	arr := strings.Split(date, " ")
	day := arr[0][:len(arr[0])-2]
	if len(day) == 1 {
		day = "0" + day
	}
	month := m[arr[1]]
	return arr[2] + "-" + month + "-" + day
}
// @lc code=end

