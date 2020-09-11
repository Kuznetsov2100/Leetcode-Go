/*
 * @lc app=leetcode.cn id=321 lang=golang
 *
 * [321] 拼接最大数
 */

// @lc code=start
// 假设从nums1数组选取x个数字，则从nums2数组选取k-x个数字， 枚举所有可能的挑选组合
// pickKdigits函数返回从nums1数组选取x个数字所能获得的最大数
// maxMergeNum函数返回能够组成的最大数
// isGreater函数判断两个数组谁更大
func maxNumber(nums1 []int, nums2 []int, k int) []int {
	var res []int
	m, n := len(nums1), len(nums2)
	for i := 0; i <= k; i++ {
		if i > m || k-i > n {
			continue
		}
		nums := maxMergeNum(pickKdigits(nums1, i), pickKdigits(nums2, k-i))
		if isGreater(nums, res) {
			res = nums // 更新res
		}
	}
	return res
}

func isGreater(numsA,numsB []int) bool {
	m, n := len(numsA), len(numsB)
	length := min(m, n)
	for i := 0; i < length; i++ {
		if numsA[i] > numsB[i] {
			return true
		} else if numsA[i] < numsB[i] {
			return false
		}
	}
	return m > n
}

func pickKdigits(nums []int, k int) []int {
	discard := len(nums)-k
	var stack []int // 单调递减栈,栈中存储的是nums数组中的数字
	for _, num := range nums {
		for discard > 0 && len(stack) > 0 && stack[len(stack)-1] < num {
			stack = stack[:len(stack)-1]
			discard--
		}
		stack = append(stack, num)
	}
	return stack[:k] // 挑选前k个数字
}

func maxMergeNum(numsA, numsB []int) []int {
	m, n := len(numsA), len(numsB)
	i, j, index := 0, 0, 0
	res := make([]int, m+n)
	for i < m && j < n {
		if isGreater(numsA[i:], numsB[j:]) {
			res[index] = numsA[i]
			i++
		} else {
			res[index] = numsB[j]
			j++
		}
		index++
	}
	for i < m {
		res[index] = numsA[i]
		i++
		index++
	}
	for j < n {
		res[index] = numsB[j]
		j++
		index++
	}
    return res
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

