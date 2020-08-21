/*
 * @lc app=leetcode.cn id=229 lang=golang
 *
 * [229] 求众数 II
 */

// @lc code=start
func majorityElement(nums []int) []int {
	// 摩尔投票法 --> 多个候选人
	var res []int
	n := len(nums)
	if n == 0 {
		return res
	}
	// 根据题意，最多只有2个元素出现的次数超过n/3次，所以设置2个候选人
	candidate1, vote1 := nums[0], 0
	candidate2, vote2 := nums[0], 0
	for _, v := range nums {
		if candidate1 == v { // 给候选人1投票
			vote1++
			continue
		}
		if candidate2 == v { // 给候选人2投票
			vote2++
			continue
		}
		if vote1 == 0 { // 候选人1的票数为0，换掉候选人1，投票给新的候选人1
			candidate1 = v
			vote1++
			continue
		}
		if vote2 == 0 { // 候选人2的票数为0，换掉候选人2，投票给新的候选人2
			candidate2 = v
			vote2++
			continue
		}
		// 无人投票给候选人1和候选人2
		vote1--
		vote2--
	}
	// 确认选出的候选人1和候选人2的票数是否大于n/3的要求
	vote1, vote2 = 0, 0
	for _, v := range nums {
		if candidate1 == v {
			vote1++
		} else if candidate2 == v {
			vote2++
		}
	}
	if vote1 > n/3 {
		res = append(res, candidate1)
	}
	if vote2 > n/3 {
		res = append(res, candidate2)
	}
	return res
}
// @lc code=end

