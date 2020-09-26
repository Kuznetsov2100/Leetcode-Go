/*
 * @lc app=leetcode.cn id=1431 lang=golang
 *
 * [1431] 拥有最多糖果的孩子
 */

// @lc code=start
func kidsWithCandies(candies []int, extraCandies int) []bool {
	maximum := math.MinInt64
	for _, num := range candies {
		if num > maximum {
			maximum = num
		}
	}
	res := make([]bool, len(candies))
	for i, num := range candies {
		if num+extraCandies >= maximum {
			res[i] = true
		}
	}
	return res
}
// @lc code=end

