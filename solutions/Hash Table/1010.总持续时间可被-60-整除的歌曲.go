/*
 * @lc app=leetcode.cn id=1010 lang=golang
 *
 * [1010] 总持续时间可被 60 整除的歌曲
 */

// @lc code=start
func numPairsDivisibleBy60(time []int) int {
	var res int
	m := make([]int, 60)
	for i := range time {
		m[time[i] % 60]++
	}
	for i, count := range m {
		if count > 0 {
			if i == 0 || i == 30 {
				res += count*(count-1)/2
			} else {
				res += count*m[60-i]
				m[60-i] = 0
			}
		}	
	}
	return res
}

// func numPairsDivisibleBy60(time []int) int {
// 	var res int
// 	m := make([]int, 60)
// 	for _, t := range time {
// 		t = t % 60
// 		if t == 0 {
// 			res += m[0]
// 		} else {
// 			res += m[60-t]
// 		}
// 		m[t]++
// 	}
// 	return res
// }
// @lc code=end

