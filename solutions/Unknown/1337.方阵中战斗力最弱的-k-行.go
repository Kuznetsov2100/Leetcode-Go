/*
 * @lc app=leetcode.cn id=1337 lang=golang
 *
 * [1337] 方阵中战斗力最弱的 K 行
 */

// @lc code=start
func kWeakestRows(mat [][]int, k int) []int {
	var arr []*foo
	for i := 0; i < len(mat); i++ {
		count := 0
		for j := 0; j < len(mat[0]); j++ {
			if mat[i][j] == 1 {
				count++
			} else {
				break
			}
		}
		arr = append(arr, &foo{index:i, value:count})
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i].value < arr[j].value {
			return true
		} else if arr[i].value == arr[j].value {
			if arr[i].index < arr[j].index {
				return true
			}
		}
		return false
	})
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = arr[i].index
	}
	return res	
}

type foo struct {
	index, value int
}
// @lc code=end

