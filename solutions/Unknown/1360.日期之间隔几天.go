/*
 * @lc app=leetcode.cn id=1360 lang=golang
 *
 * [1360] 日期之间隔几天
 */

// @lc code=start
func daysBetweenDates(date1 string, date2 string) int {
	date1arr := strings.Split(date1, "-")
	date2arr := strings.Split(date2, "-")
	var arr1, arr2 []int
	for i := 0; i < len(date1arr); i++ {
		number1, _ := strconv.Atoi(date1arr[i])
		number2, _ := strconv.Atoi(date2arr[i])
		arr1 = append(arr1, number1)
		arr2 = append(arr2, number2)
	}
	
	start := time.Date(arr1[0],time.Month(arr1[1]), arr1[2], 0, 0, 0, 0, time.UTC)
	end := time.Date(arr2[0], time.Month(arr2[1]), arr2[2], 0, 0, 0, 0, time.UTC)
	return abs(int(end.Sub(start)/86400000000000))
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

