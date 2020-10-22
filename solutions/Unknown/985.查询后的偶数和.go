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
	for i := range queries {
		val, index := queries[i][0], queries[i][1]
		if A[index] % 2 == 0 {
			evenSum -= A[index]
		}
		A[index] += val
		if A[index] % 2 == 0 {
			evenSum += A[index]
		}
		res[i] = evenSum
	}
	return res
}
// @lc code=end

