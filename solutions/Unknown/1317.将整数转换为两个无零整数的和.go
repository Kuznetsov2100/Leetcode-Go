/*
 * @lc app=leetcode.cn id=1317 lang=golang
 *
 * [1317] 将整数转换为两个无零整数的和
 */

// @lc code=start
func getNoZeroIntegers(n int) []int {
	var res []int
	for i := 1; i < n; i++ {
		a := strconv.Itoa(i)
		b := strconv.Itoa(n-i)
		if !strings.Contains(a, "0") && !strings.Contains(b, "0") {
			res = []int{i, n-i}
			break
		}
	}
	return res
}
// @lc code=end

