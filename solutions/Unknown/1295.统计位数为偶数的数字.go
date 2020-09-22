/*
 * @lc app=leetcode.cn id=1295 lang=golang
 *
 * [1295] 统计位数为偶数的数字
 */

// @lc code=start
func findNumbers(nums []int) int {
	var res int
	for _, num := range nums {
		if count(num) % 2 == 0 {
			res++
		}
	}
	return res
}

func count(num int) int {
	var res int
	for num > 0 {
		num /= 10
		res++
	}
	return res
}
// @lc code=end

