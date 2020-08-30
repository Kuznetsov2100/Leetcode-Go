/*
 * @lc app=leetcode.cn id=165 lang=golang
 *
 * [165] 比较版本号
 */

// @lc code=start
func compareVersion(version1 string, version2 string) int {
	str_arr1 := strings.Split(version1, ".")
	str_arr2 := strings.Split(version2, ".")
	m, n := len(str_arr1), len(str_arr2)
	if m > n {
		for i := m-n; i > 0; i-- {
			str_arr2 = append(str_arr2, "0")
		}
	} else if m < n {
		for i := n-m; i > 0; i-- {
			str_arr1 = append(str_arr1, "0")
		}
	}
	for i := 0; i < len(str_arr1); i++ {
		number1, _ := strconv.Atoi(str_arr1[i])
		number2, _ := strconv.Atoi(str_arr2[i])
		if number1 < number2 {
			return -1
		} else if number1 > number2 {
			return 1
		}
	}
	return 0
}
// @lc code=end

