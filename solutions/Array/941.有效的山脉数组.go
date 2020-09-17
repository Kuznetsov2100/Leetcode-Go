/*
 * @lc app=leetcode.cn id=941 lang=golang
 *
 * [941] 有效的山脉数组
 */

// @lc code=start
// 双指针
func validMountainArray(A []int) bool {
	n := len(A)
	if n < 3 {
		return false
	}
	i, j := 0, n-1
	for i  < j-1 {
		if A[i+1] <= A[i] && A[j-1] <= A[j] {
			return false
		}
		if A[i] < A[i+1] {
			i++
		}
		if A[j-1] > A[j] {
			j--
		}
	}
	if i == 0 || j == n-1 || j-i > 1 {
		return false
	}
	return true
}

// 线性扫描
// func validMountainArray(A []int) bool {
// 	n := len(A)
// 	i := 0
// 	// walk up
// 	for i < n-1 && A[i] < A[i+1] {
// 		i++
// 	}
// 	if i == 0 || i == n-1 {
// 		return false
// 	}
// 	// walk down
// 	for i < n-1 && A[i] > A[i+1] {
// 		i++
// 	}
// 	return i == n-1
// }
// @lc code=end

