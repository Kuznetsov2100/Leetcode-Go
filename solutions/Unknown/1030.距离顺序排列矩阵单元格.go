/*
 * @lc app=leetcode.cn id=1030 lang=golang
 *
 * [1030] 距离顺序排列矩阵单元格
 */

// @lc code=start
func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	var res [][]int
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			res = append(res, []int{i, j})
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return (abs(res[i][0]-r0)+abs(res[i][1]-c0)) < (abs(res[j][0]-r0)+abs(res[j][1]-c0))
	})
	return res
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
// @lc code=end

