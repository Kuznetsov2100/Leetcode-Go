/*
 * @lc app=leetcode.cn id=914 lang=golang
 *
 * [914] 卡牌分组
 */

// @lc code=start
func hasGroupsSizeX(deck []int) bool {
	// 统计各个数的次数， 然后求这些次数之间的最大公约数
	if len(deck) < 2 {
		return false
	}
	m := make(map[int]int)
	for i := range deck {
		m[deck[i]]++
	}
	commonDivisor := m[deck[0]]
	for _, count := range m {
		commonDivisor = gcd(commonDivisor, count)
		if commonDivisor < 2 {
			return false
		}
	}
	return true
}

func gcd(a, b int) int {
	if a == 0 {
		return b
	}
	return gcd(b % a, a)
}
// @lc code=end

