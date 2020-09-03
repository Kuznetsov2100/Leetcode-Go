/*
 * @lc app=leetcode.cn id=75 lang=golang
 *
 * [75] 颜色分类
 */

// @lc code=start
func sortColors(nums []int)  {
	// 三路快排思想
	// 它从左到右遍历数组一次，维护一个指针lt使得a[lo..lt-1]中的元素都小于pivot，
	// 一个指针gt使得a[gt+1..hi]中的元素都大于pivot，
	// 一个指针i使得a[lt..i-1] 中的元素都等于pivot，a[i..gt]中的元素都还未确定，
	pivot := 1
	lt := 0
	gt := len(nums)-1
	i := 0
	for i <= gt {
		if nums[i] < pivot {
			nums[i], nums[lt] = nums[lt], nums[i]
			lt++
			i++ // 同理，交换后，nums[i]也就是原来的nums[lt]已经比较过,状态已经确定，所以指针i需要跳过这一个元素到后面去
		} else if nums[i] > pivot {
			nums[i], nums[gt] = nums[gt], nums[i]
			gt-- 
			// 为啥这里不需要i++, 因为nums[i]与nums[gt]交换后，原来的nums[gt]也就是现在的nums[i]还没有进行比较，
			// 处于未确定状态，如果i++，这个元素就被遗漏了
		} else {
			i++
		}
	}
}
// @lc code=end

