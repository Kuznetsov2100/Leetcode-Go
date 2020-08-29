/*
 * @lc app=leetcode.cn id=74 lang=golang
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}
	m, n := len(matrix), len(matrix[0])
	//  row := mid/n, col := mid%n
	// 将一维数组坐标映射到矩阵中的相应坐标,就可以对矩阵直接进行二分搜索
	left, right := 0, m * n -1
	for left <= right {
		mid := left + (right - left)/2
		if val := matrix[mid/n][mid%n]; val  < target {
			left = mid + 1
		} else if val > target {
			right = mid - 1
		} else if val == target {
			return true
		}
	}
	return false
}

// 个人解法： 先利用二分搜索找到行坐标，再在矩阵相应行执行二分搜索
// func searchMatrix(matrix [][]int, target int) bool {
// 	if len(matrix) == 0 || len(matrix[0]) == 0 {
// 		return false
// 	}
// 	m, n := len(matrix), len(matrix[0])
// 	left, right := 0, m-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if matrix[mid][0] <= target && target <= matrix[mid][n-1] {
// 			return binarySearch(matrix[mid],target)			
// 		} else if target < matrix[mid][0] {
// 			right = mid - 1
// 		} else if target > matrix[mid][0] {
// 			left = mid + 1
// 		}
// 	}
// 	return false
// }

// func binarySearch(arr []int, target int) bool {
// 	left, right := 0, len(arr)-1
// 	for left <= right {
// 		mid := left + (right - left)/2
// 		if arr[mid] == target {
// 			return true
// 		} else if arr[mid] < target {
// 			left = mid + 1
// 		} else if arr[mid] > target {
// 			right = mid - 1
// 		}
// 	}
// 	return false
// }
// @lc code=end

