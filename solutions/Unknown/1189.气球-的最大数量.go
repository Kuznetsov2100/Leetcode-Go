/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	m := make([]int, 26)
	for i := range text {
		m[text[i]-'a']++
	}
	banCount, loCount := math.MaxInt64, math.MaxInt64
	for i, count := range m {
		alpha := i+'a'
		if alpha == 'b'  || alpha == 'a' || alpha == 'n' {
			banCount = min(banCount, count)
		} else if alpha == 'l' || alpha == 'o' {
			loCount = min(loCount, count)
		}
	}
	return min(banCount, loCount/2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

