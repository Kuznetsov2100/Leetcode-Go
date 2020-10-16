/*
 * @lc app=leetcode.cn id=836 lang=golang
 *
 * [836] 矩形重叠
 */

// @lc code=start
// 版本一：检查重叠区域面积是否不为0
func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	return commmonArea(rec1[0], rec1[1], rec1[2], rec1[3], rec2[0], rec2[1], rec2[2], rec2[3]) != 0
}

func commmonArea(A, B, C, D, E, F, G, H int) int {
	x1 := max(A, E)
	y1 := max(B, F)
	x2 := min(C, G)
	y2 := min(D, H)
	if x1 < x2 && y1 < y2 {
		return (x2-x1)*(y2-y1)
	}
	return 0
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

// 版本二：排除掉不重叠的情况，剩下的即为重叠的情况
// func isRectangleOverlap(rec1 []int, rec2 []int) bool {
// 	return !(rec1[0] >= rec2[2] || rec1[2] <= rec2[0] || rec1[3] <= rec2[1] || rec1[1] >= rec2[3])
// }
// @lc code=end

