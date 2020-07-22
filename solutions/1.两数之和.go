/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
    record := make(map[int]int)
	for index, value := range nums {
		complement := target - value
		if _, ok := record[complement]; ok {
			return []int{record[complement], index}
		} else {
			record[value] = index
		}
	}
	return nil
}
// @lc code=end

