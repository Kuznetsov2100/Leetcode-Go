/*
 * @lc app=leetcode.cn id=134 lang=golang
 *
 * [134] 加油站
 */

// @lc code=start
func canCompleteCircuit(gas []int, cost []int) int {
	res := 0
	n := len(gas)
	total_rest, cur_rest := 0, 0 
	for i := 0; i < n; i++ {
		total_rest += gas[i]-cost[i]
		cur_rest   += gas[i]-cost[i]
		if cur_rest < 0 {
			res = i + 1
			cur_rest = 0
		}
	}
	if total_rest < 0 {
		return -1
	}
	return res
}


// 暴力解法
// func canCompleteCircuit(gas []int, cost []int) int {
// 	res := -1
// 	n := len(gas)
// 	for i := 0; i < n; i++ {
// 		if cost[i] <= gas[i] {
// 			curgas := gas[i]
// 			j := i
// 			k := 0
// 			for ; k < n; k++ {
// 				curgas -= cost[j]
// 				if curgas < 0 {
// 					break
// 				}
// 				j = (j+1) % n
// 				curgas += gas[j]
// 			}
// 			if k == n {
// 				res = i
// 				return res
// 			}
// 		}	
// 	}
// 	return res
// }


// @lc code=end

