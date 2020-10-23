/*
 * @lc app=leetcode.cn id=1237 lang=golang
 *
 * [1237] 找出给定方程的正整数解
 */

// @lc code=start
/** 
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	x, y := 1, 1000
	var res [][]int
	for x <= 1000 && y >= 1 {
		ans := customFunction(x, y)
		if ans < z {
			x++
		} else if ans > z {
			y--
		} else if ans == z {
			res = append(res, []int{x, y})
			x++
			y--
		}
	}
	return res
}
// @lc code=end

