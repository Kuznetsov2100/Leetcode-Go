/*
 * @lc app=leetcode.cn id=1314 lang=golang
 *
 * [1314] 矩阵区域和
 */

// @lc code=start
func matrixBlockSum(mat [][]int, K int) [][]int {
	// 本题可参考第304题， 将条件i - K <= r <= i + K, j - K <= c <= j + K 转化为row1,col1,row2,col2
	// (row1, col1)为answer[i,j]区域和的左上角坐标，(row2, col2)为answer[i,j]区域和的右下角坐标
	
	m, n := len(mat), len(mat[0])
	res := make([][]int, m)
	for i := range res {
		res[i] = make([]int, n)
	}

	dp := make([][]int, m+1) // dp数组在matrix矩阵基础上增加了一行0，一列0，可避免很多边界条件判断
	for i := range dp {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dp[i+1][j+1] = dp[i][j+1] + dp[i+1][j] - dp[i][j] + mat[i][j]
		}
	}

	sumRegion := func(i, j int) int {
		row1 := max(0, i-K)
		col1 := max(0, j-K)
		row2 := min(m-1, i+K)
		col2 := min(n-1, j+K)
		return dp[row2+1][col2+1]-dp[row1][col2+1]-dp[row2+1][col1]+dp[row1][col1]	
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			res[i][j] = sumRegion(i, j)
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

