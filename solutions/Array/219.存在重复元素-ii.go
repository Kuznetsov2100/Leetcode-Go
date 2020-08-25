/*
 * @lc app=leetcode.cn id=219 lang=golang
 *
 * [219] 存在重复元素 II
 */

// @lc code=start
func containsNearbyDuplicate(nums []int, k int) bool {
	// 维护一个哈希表，里面始终最多包含 k 个元素，当出现重复值时则说明在 k 距离内存在重复元素
	// 每次遍历一个元素则将其加入哈希表中，如果哈希表的大小大于 k，则移除哈希表中最早进入的元素
	m := make(map[int]int)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = 1
		if len(m) > k {
			delete(m, nums[i-k])
		}
	}
	return false
}
// @lc code=end

