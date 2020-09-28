/*
 * @lc app=leetcode.cn id=1323 lang=golang
 *
 * [1323] 6 和 9 组成的最大数字
 */

// @lc code=start
func maximum69Number (num int) int {
	numbyte := []byte(strconv.Itoa(num))
	for i := range numbyte {
		if numbyte[i] == '6' {
			numbyte[i] = '9'
			break
		}
	}
	number, _ := strconv.Atoi(string(numbyte))
	return number
}
// @lc code=end

