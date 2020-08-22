/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 */

// @lc code=start
func findRelativeRanks(nums []int) []string {
	var res []string
	arr := make([]int, len(nums))
	copy(arr, nums)
	sort.Slice(arr, func(i, j int)bool {return arr[i] > arr[j] })
	m := make(map[int]int)
	for i, v := range arr {
		m[v] = i+1
	}
	for _ , v := range nums {
		if rank, _ := m[v]; rank == 1 {
			res = append(res, "Gold Medal")
		} else if rank == 2 {
			res = append(res, "Silver Medal")
		} else if rank == 3 {
			res = append(res, "Bronze Medal")
		} else {
			res = append(res, strconv.Itoa(rank))
		}
	}
	return res
}
// @lc code=end

