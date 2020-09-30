/*
 * @lc app=leetcode.cn id=1360 lang=golang
 *
 * [1360] 日期之间隔几天
 */

// @lc code=start
func daysBetweenDates(date1 string, date2 string) int {
	return abs(calDays(date1)-calDays(date2))
}

func calDays(date string) int {
	// anchor : 1971-01-01
	var res int
	months := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	arr := strings.Split(date, "-")
	year, _ := strconv.Atoi(arr[0])
	month, _ := strconv.Atoi(arr[1])
	day, _ := strconv.Atoi(arr[2])

	// year diff
	for i := 1971; i < year; i++ {
		if isLeapYear(i) {
			res += 366
		} else {
			res += 365
		}
	}

	// month diff
	for i := 1; i < month; i++ {
		if i == 2 {
			if isLeapYear(year) {
				res += 29
			} else {
				res += 28
			}
			continue
		}
		res += months[i-1]
	}
	res += day-1 // day diff
	return res
}

func isLeapYear(year int) bool {
	if year % 100 == 0 {
		if year % 400 == 0 {
			return true
		}
	} else if year % 100 != 0 {
		if year % 4 == 0 {
			return true
		}
	}
	return false
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

