/*
 * @lc app=leetcode.cn id=762 lang=golang
 *
 * [762] 二进制表示中质数个计算置位
 */

// @lc code=start
func countPrimeSetBits(L int, R int) int {
	var res int
	for i := L; i <= R; i++ {
		count1 := hammingWeight(i)
		res += (665772 >> count1) & 1 // 665772 --> binary format: 10100010100010101100 
	}
	return res
}

func hammingWeight(num int) int {
	var res int
	for num != 0 {
		num = num & (num-1) // n & (n-1)消除数字n的二进制表示中的最后一个1
		res++
	}
	return res
}
// @lc code=end

