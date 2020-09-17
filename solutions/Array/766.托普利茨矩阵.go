/*
 * @lc app=leetcode.cn id=766 lang=golang
 *
 * [766] 托普利茨矩阵
 */

// @lc code=start
func isToeplitzMatrix(matrix [][]int) bool {
	// 满足托普利茨矩阵定义的矩阵，从第二行开始，每个元素等于其左上方的元素(跳过第一列)
	m, n := len(matrix), len(matrix[0])
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] != matrix[i-1][j-1] {
				return false
			}
		}
	}
	return true
}


// func isToeplitzMatrix(matrix [][]int) bool {
// 	m, n := len(matrix), len(matrix[0])

// 	for j := n-1; j >= 0; j-- {
// 		pivot := 100
// 		for k := 0; k+j < n && k < m; k++ {
// 			if k == 0 {
// 				pivot = matrix[k][k+j]
// 			} else {
// 				if matrix[k][k+j] != pivot {
// 					return false
// 				}
// 			}
// 		}
// 	}
	
// 	for i := 1; i < m; i++ {
// 		pivot := 100
// 		for k := 0; k+i < m && k < n; k++ {
// 			if k == 0 {
// 				pivot = matrix[i+k][k]
// 			} else {
// 				if matrix[i+k][k] != pivot {
// 					return false
// 				}
// 			}
// 		}
// 	}
// 	return true
// }
// @lc code=end

