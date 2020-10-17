/*
 * @lc app=leetcode.cn id=832 lang=golang
 *
 * [832] 翻转图像
 */

// @lc code=start
func flipAndInvertImage(A [][]int) [][]int {
	m, n := len(A), len(A[0])
	for i := 0; i < m; i++ {
		for j , k := 0, n-1; j <= k; j, k= j+1, k-1 {
			A[i][j], A[i][k] = 1 ^ A[i][k], 1 ^ A[i][j] // 异或运算， 注意循环结束条件为:j <= k
		}
	}	
	return A
}
// @lc code=end

