/*
 * @lc app=leetcode.cn id=860 lang=golang
 *
 * [860] 柠檬水找零
 */

// @lc code=start
func lemonadeChange(bills []int) bool {
	var five, ten int
	for _, bill := range bills {
		if bill == 5 {
			five++
		} else if bill == 10 {
			if five == 0 {
				return false
			}
			five--
			ten++
		} else {
			if five == 0 {
				return false
			}
			if ten == 0 {
				if five < 3 {
					return false
				}
				five -= 3
			} else {
				ten--
				five--
			}
		}
	}
	return true
}
// @lc code=end

