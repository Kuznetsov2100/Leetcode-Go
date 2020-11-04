/*
 * @lc app=leetcode.cn id=846 lang=golang
 *
 * [846] 一手顺子
 */

// @lc code=start
func isNStraightHand(hand []int, W int) bool {
	if len(hand) % W != 0 {
		return false
	}
	sort.Ints(hand)
	m := make(map[int]int)
	for i := range hand {
		m[hand[i]]++
	}
	for i := range hand {
		if m[hand[i]] > 0 {
			for j := 0; j < W; j++ {
				m[hand[i]+j]--
				if m[hand[i]+j] < 0 {
					return false
				}
			}
		}
	}
	return true	
}
// @lc code=end

