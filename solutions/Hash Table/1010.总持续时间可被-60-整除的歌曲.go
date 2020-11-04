/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	var res int
	remainderMap := make([]int, 60)
	for _, t := range time {
		t = t % 60
		if t == 0 {
			res += remainderMap[0]
		} else {
			res += remainderMap[60-t]
		}
		remainderMap[t]++
	}
	return res
}
// @lc code=end

