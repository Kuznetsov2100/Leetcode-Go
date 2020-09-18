/*
 * @lc app=leetcode.cn id=1128 lang=golang
 *
 * [1128] 等价多米诺骨牌对的数量
 */

// @lc code=start
func numEquivDominoPairs(dominoes [][]int) int {
	// 将一张牌较小的数字作为十位，较大的数字作为各位，将它映射为一个两位数，
	// 如果两张牌等价，则它们会映射到相同的数字
	// 然后遍历key数组，根据n选2的组合公式 n*(n-1)/2，计算答案。
	var res int
	key := make([]int, 100)
	for _, domino := range dominoes {
		var decimal int
		if domino[0] < domino[1] {
			decimal = domino[0]*10+domino[1]
		} else {
			decimal = domino[1]*10+domino[0]
		}
		key[decimal]++
	}
	// 因为映射的数字范围是[11,99]
	for i := 11; i < 100; i++ {
		if key[i] > 0 {
			res += key[i]*(key[i]-1)/2
		}
	}
	return res
}
// @lc code=end

