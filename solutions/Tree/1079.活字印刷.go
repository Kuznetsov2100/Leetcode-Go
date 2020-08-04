/*
 * @lc app=leetcode.cn id=1079 lang=golang
 *
 * [1079] 活字印刷
 */

// @lc code=start
import "sort"
func numTilePossibilities(tiles string) int {
	nums := make([]int, len(tiles))
	for i := range tiles {
		nums[i] = int(tiles[i])
	}

	// 求出含有重复数字的所有可能的子集
	sort.Ints(nums)
	var res [][]int
	used := make([]bool, len(nums))
	var backtrack func(start int, track []int)
	backtrack = func(start int, track []int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		res = append(res, tmp)
		for i := start; i < len(nums); i++ {
			if !used[i] {
				// 剪枝
				if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
					continue
				}
				track = append(track, nums[i])
				used[i] = true
				backtrack(i+1, track)
				track = track[:len(track)-1]
				used[i] = false
			}
		}
	}
	backtrack(0, []int{})
	// 去掉第一个空集
	res = res[1:]
	sum := 0
	// 对每一个子集(含有重复数字的数组)求出不重复的所有全排列个数，然后加总
	for i := range res {
		sum += countPermuteUnique(res[i])
	}
	return sum
}

func countPermuteUnique(nums []int) int {
	sort.Ints(nums)
	var res int
	used := make([]bool, len(nums))
	var bt func(path []int)
	bt = func(path []int) {
		if len(path) == len(nums) {
			res++
			return
		}
		for i := range nums {
			if !used[i] { // 访问未访问的
				if i > 0 && nums[i-1] == nums[i] && !used[i-1] {
					continue
				}
				path = append(path, nums[i]) // 做选择
				used[i] = true
				bt(path) // 进入下一层决策树
				path = path[:len(path)-1] // 撤销选择
				used[i] = false
			}
		}
	}
	bt([]int{})
	return res
}
// @lc code=end