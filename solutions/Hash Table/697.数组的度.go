/*
 * @lc app=leetcode.cn id=697 lang=golang
 *
 * [697] 数组的度
 */

// @lc code=start
func findShortestSubArray(nums []int) int {
	res, maxFrequency := 50001, 0
	m := make(map[int]*Number)
	for i, num := range nums {
		if _, ok := m[num]; !ok {
			m[num] = &Number{count: 1, firstIndex: i, lastIndex: i}
		} else {
			m[num].count++
			m[num].lastIndex = i
		}
		if c := m[num].count; c > maxFrequency {
			maxFrequency = c
		}
	}
	for _, number := range m {
		if number.count == maxFrequency {
			res = min(res, number.lastIndex-number.firstIndex+1)
		}
	}
	return res
}

type Number struct {
	count int
	firstIndex int
	lastIndex int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

