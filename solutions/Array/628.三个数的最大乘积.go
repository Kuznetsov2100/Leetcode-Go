/*
 * @lc app=leetcode.cn id=628 lang=golang
 *
 * [628] 三个数的最大乘积
 */

// @lc code=start
func maximumProduct(nums []int) int {
	first, second,third := math.MinInt64, math.MinInt64, math.MinInt64
	minimum, secondToMin := math.MaxInt64, math.MaxInt64
	for _, v := range nums {
		if v > first {
			first, second, third = v, first, second
		} else if v > second {
			second, third = v, second
		} else if v > third {
			third = v
		}
		if v < minimum {
			minimum, secondToMin = v, minimum
		} else if v < secondToMin {
			secondToMin = v
		}
	}
	return max(first*second*third, first*minimum*secondToMin)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
// @lc code=end

