/*
 * @lc app=leetcode.cn id=905 lang=golang
 *
 * [905] 按奇偶排序数组
 */

// @lc code=start
func sortArrayByParity(A []int) []int {
	var even,odd []int
	for _, num := range A {
		if num % 2 == 0 {
			even = append(even, num)
		} else {
			odd = append(odd, num)
		}
	}
	return append(even, odd...)
}
// @lc code=end

