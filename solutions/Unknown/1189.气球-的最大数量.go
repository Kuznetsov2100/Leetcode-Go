/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	var b,a,n,l,o int
	for i := range text {
		switch text[i] {
			case 'b':
				b++
			case 'a':
				a++
			case 'n':
				n++
			case 'l':
				l++
			case 'o':
				o++
		}
	}
	return min(min(min(b, a), n),min(l, o)/2)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

