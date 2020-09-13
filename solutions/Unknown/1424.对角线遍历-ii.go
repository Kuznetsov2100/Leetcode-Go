/*
 * @lc app=leetcode.cn id=1424 lang=golang
 *
 * [1424] 对角线遍历 II
 */

// @lc code=start
func findDiagonalOrder(nums [][]int) []int {
	//	横纵坐标之和相等时，根据横坐标降序排序
	//  横纵坐标之和不相等时，根据坐标和升序排序
	var res []int
	type point struct {
		row,col,value int
	}
	var arr []*point
	for i := range nums {
		for j := range nums[i] {
			arr = append(arr, &point{row:i, col:j, value: nums[i][j]})
		}
	}
	sort.Slice(arr, func(i, j int) bool {
		sumA, sumB := arr[i].row+arr[i].col, arr[j].row+arr[j].col
		if sumA == sumB {
			if arr[i].row > arr[j].row {
				return true
			}
		} else if sumA < sumB {
			return true
		}
		return false
	})
	for i := range arr {
		res = append(res, arr[i].value)
	}
	return res
}
// @lc code=end

