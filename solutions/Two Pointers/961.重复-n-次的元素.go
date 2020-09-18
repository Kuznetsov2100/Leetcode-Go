/*
 * @lc app=leetcode.cn id=961 lang=golang
 *
 * [961] 重复 N 次的元素
 */

// @lc code=start
func repeatedNTimes(A []int) int {
	var res int
	sort.Ints(A)
	i, j := 0, len(A)-1
	for i < j {
		if A[i] < A[i+1] {
			i++
		} else {
			res = A[i]
			break
		}
		if A[j-1] < A[j] {
			j--
		} else {
			res = A[j]
			break
		}
	}
	return res
}
// @lc code=end

