/*
 * @lc app=leetcode.cn id=1013 lang=golang
 *
 * [1013] 将数组分成和相等的三个部分
 */

// @lc code=start
func canThreePartsEqualSum(A []int) bool {
	var sum int
	for i := range A {
		sum += A[i]
	}
	if sum % 3 != 0 {
		return false
	}
	sum /= 3
	count , curSum := 0, 0
	for i := 0; i < len(A)-1; i++ {
		curSum += A[i]
		if curSum == sum {
			count++
			if count == 2 {
				return true
			}
			curSum = 0
		}
	}
	return false
}
// @lc code=end

