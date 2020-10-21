/*
 * @lc app=leetcode.cn id=985 lang=golang
 *
 * [985] 查询后的偶数和
 */

// @lc code=start
func sumEvenAfterQueries(A []int, queries [][]int) []int {
	res := make([]int, len(queries))
	var evenSum int
	for _, num := range A {
		if num % 2 == 0 {
			evenSum += num
		}
	}
	for i, q := range queries {
		keep := A[q[1]]
		A[q[1]] += q[0]
		if keep % 2 == 0 {
			if q[0] % 2 ==  0 {
				evenSum += q[0]
			} else {
				evenSum -= keep
			}
		} else {
			if q[0] % 2 != 0 {
				evenSum += A[q[1]]
			}
		}
		res[i] = evenSum
	}
	return res
}
// @lc code=end

