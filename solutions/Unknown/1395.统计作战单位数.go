/*
 * @lc app=leetcode.cn id=1395 lang=golang
 *
 * [1395] 统计作战单位数
 */

// @lc code=start
func numTeams(rating []int) int {
	// 假设三元组为(x,y,z)，遍历rating数组，枚举每一个可能的y,统计y左边比y小的评分数量leftsmaller,比y大的评分数量leftgreater
	// 统计y右边比y小的评分数量rightsmaller,比y大的评分数量rightgreater
	// 以y为中间点的作战单位总数 = leftsmaller*rightgreater+leftgreater*rightsmaller
	// 遍历完成即可得到总的作战单位数
	var res int
	n := len(rating)
	for i := 1; i < n-1; i++ {
		leftgreater, leftsmaller := 0, 0
		for j := 0; j < i; j++ {
			if rating[j] > rating[i] {
				leftgreater++
			} else if rating[j] < rating[i] {
				leftsmaller++
			}
		}
		rightgreater, rightsmaller := 0, 0
		for k := i+1; k < n; k++ {
			if rating[k] > rating[i] {
				rightgreater++
			} else if rating[k] < rating[i] {
				rightsmaller++
			}
		}
		res += leftsmaller*rightgreater+leftgreater*rightsmaller
	}
	return res
}
// @lc code=end

