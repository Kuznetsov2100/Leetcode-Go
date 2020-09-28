/*
 * @lc app=leetcode.cn id=1385 lang=golang
 *
 * [1385] 两个数组间的距离值
 */

// @lc code=start
func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	var res int
	n := len(arr2)
	sort.Ints(arr2)
	for _, num := range arr1 {
		index := sort.SearchInts(arr2, num)
		var diff int
		if index == 0 {
			diff = abs(arr2[index]-num)
		} else if index == n {
			diff = abs(arr2[index-1]-num)
		} else {
			diff = min(abs(arr2[index-1]-num), abs(arr2[index]-num))
		}
		if diff > d {
			res++
		}
	}
	return res
}



func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

