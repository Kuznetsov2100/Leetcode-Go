/*
 * @lc app=leetcode.cn id=679 lang=golang
 *
 * [679] 24 点游戏
 */

// @lc code=start
func judgePoint24(nums []int) bool {
	for  _, arr := range permute(nums) {
		if judge4(float64(arr[0]), float64(arr[1]), float64(arr[2]), float64(arr[3])) {
			return true
		}
	}
	return false
}

// a,b,c,d 是全排列中的某一个排列，a,b,c,d无需考虑顺序
func judge4(a,b,c,d float64) bool {
	return judge3(a+b, c, d)  ||
		   judge3(a-b, c, d)  ||
		   judge3(a*b, c, d)  ||
		   judge3(a/b, c, d)
}

// a是两数运算的结果，a与b,c中的任何一个都不等价， b和c等价
func judge3(a, b, c float64) bool {
	return judge2(a+b, c) ||
		   judge2(a-b, c) ||
		   judge2(a*b, c) ||
		   judge2(a/b, c) ||
		   judge2(b-a, c) ||
		   judge2(b/a, c) ||
		   judge2(a, b+c) ||
		   judge2(a, b-c) ||
		   judge2(a, b*c) ||
		   judge2(a, b/c)
}

// a 和 b不等价
func judge2(a, b float64) bool {
	const epslion = 1e-3
	return math.Abs(a+b-24) < epslion ||
		   math.Abs(a-b-24) < epslion ||
		   math.Abs(a*b-24) < epslion ||
		   math.Abs(a/b-24) < epslion ||
		   math.Abs(b-a-24) < epslion ||
		   math.Abs(b/a-24) < epslion
}

func permute(nums []int) [][]int {
	var res [][]int
	used := make([]bool, len(nums))
	var bt func(path []int)
	bt = func(path []int) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp,path)
			res = append(res, tmp)
			return
		}
		for i := range nums {
			if !used[i] { // 访问未访问的
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

