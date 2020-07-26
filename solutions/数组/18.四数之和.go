/*
 * @lc app=leetcode.cn id=18 lang=golang
 *
 * [18] 四数之和
 */

// @lc code=start
import "sort"
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSumTarget(nums,4,0,target)
}
// 调用nSumTarget之前需要给nums排序
func nSumTarget(nums []int, n, start, target int) (res [][]int) {
	length := len(nums)
	if n < 2 || n > length { // 至少是2sum, nums长度不应小于n
		return res
	}
	if n == 2 { // base case, 2sum
		lo ,hi := start, length-1
		for lo < hi { // 双指针
			left := nums[lo]
			right := nums[hi]
			sum := left + right	
			if sum < target { // 左指针向右移动，使sum变大
				for lo < hi && nums[lo] == left { // 跳过所有相邻的重复元素
					lo++
				}
			} else if sum > target { // 右指针向左移动,使sum变小
				for lo < hi && nums[hi] == right { // 跳过所有相邻的重复元素
					hi--
				}
			} else { 
				res = append(res, []int{left,right}) // 添加符合要求的元素对
				for lo < hi && nums[lo] == left { // 跳过所有相邻的重复元素
					lo++
				}
				for lo < hi && nums[hi] == right { // 跳过所有相邻的重复元素
					hi--
				}
			}
		}
	} else {
		for i := start; i < length; i++ {
			sub := nSumTarget(nums,n-1,i+1,target-nums[i])
			for _, arr := range sub {
				arr = append([]int{nums[i]},arr...) // 拼接 nums[i]和n-1Sum的结果组成nSum的结果
				res = append(res,arr) // 添加符合要求的结果
			}
			for i < length-1 && nums[i] == nums[i+1] { // 跳过和第一个元素重复的所有元素
				i++
			}
		}
	}
	return res
}
// @lc code=end

