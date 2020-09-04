/*
 * @lc app=leetcode.cn id=304 lang=golang
 *
 * [304] 二维区域和检索 - 矩阵不可变
 */

// @lc code=start
type NumMatrix struct {
	dp [][]int
}


func Constructor(matrix [][]int) NumMatrix {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return NumMatrix{}
	}
	m, n := len(matrix), len(matrix[0])
	dp := make([][]int, m+1) // dp数组在matrix矩阵基础上增加了一行0，一列0，可避免很多边界条件判断
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + matrix[i][j]
		}
	}
	return NumMatrix{dp:dp}
}

// sum(ABCD) = sum(OD)-sum(OB)-sum(OC)+sum(OA)
func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	return this.dp[row2+1][col2+1]-this.dp[row1][col2+1]-this.dp[row2+1][col1]+this.dp[row1][col1]
}


/**
 * Your NumMatrix object will be instantiated and called as such:
 * obj := Constructor(matrix);
 * param_1 := obj.SumRegion(row1,col1,row2,col2);
 */
// @lc code=end

