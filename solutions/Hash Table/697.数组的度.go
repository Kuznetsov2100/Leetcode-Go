/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	res, maxFrequency := 50001, 1
	m := make(map[int]*Number)
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = &Number{count: 1, firstIndex: i, lastIndex: i}
		} else {
			m[num].count++
			m[num].lastIndex = i
			if c := m[num].count; c > maxFrequency {
				maxFrequency = c
			}
		}
	}
	for _, number := range m {
		if number.count == maxFrequency {
			length := number.lastIndex-number.firstIndex+1
			if length < res {
				res = length
			}
		}
	}
	return res
}

type Number struct {
	count, firstIndex, lastIndex int
}
// @lc code=end

