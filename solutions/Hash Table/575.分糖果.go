/*
 * @lc app=leetcode.cn id=575 lang=golang
 *
 * [575] 分糖果
 */

// @lc code=start
func distributeCandies(candies []int) int {
	m := make(map[int]bool)
	for i := range candies {
		m[candies[i]] = true
	}
	return min(len(m), len(candies)/2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

