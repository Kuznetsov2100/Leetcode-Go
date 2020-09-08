/*
 * @lc app=leetcode.cn id=223 lang=golang
 *
 * [223] 矩形面积
 */

// @lc code=start
func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	area1 := (C-A)*(D-B)
	area2 := (G-E)*(H-F)
	return area1+area2-commmonArea(A,B,C,D,E,F,G,H)
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
// @lc code=end

