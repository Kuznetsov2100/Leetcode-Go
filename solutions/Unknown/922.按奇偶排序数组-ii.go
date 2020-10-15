/*
 * @lc app=leetcode.cn id=922 lang=golang
 *
 * [922] 按奇偶排序数组 II
 */

// @lc code=start
func sortArrayByParityII(A []int) []int {
	res := make([]int, len(A))
	evenIndex, oddIndex := 0, 1
	for i := range A {
		if A[i] % 2 == 0 {
			res[evenIndex] = A[i]
			evenIndex += 2
		} else {
			res[oddIndex] = A[i]
			oddIndex += 2
		}
	}
	return res
}
// @lc code=end

