/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */

// @lc code=start
func sortedSquares(A []int) []int {
	// 数组A按非递减顺序排序，那么数组元素的平方的最大值只可能在左右两端,不可能出现在数组的中间。
	// 使用双指针指向左右两端，每次判断两个指针指向的数的绝对值,
	// 优先选择绝对值大的数来生成平方数，逆序填入新数组
	n := len(A)
	res := make([]int, n)
	i, j := 0, n-1
	for pos := n-1; pos >= 0; pos-- {
		if A[i] + A[j] < 0 {
			res[pos] = A[i]*A[i]
			i++
		} else {
			res[pos] = A[j]*A[j]
			j--
		}
	}
	return res
}
// @lc code=end

