/*
 * @lc app=leetcode.cn id=690 lang=golang
 *
 * [690] 员工的重要性
 */

// @lc code=start
/**
 * Definition for Employee.
 * type Employee struct {
 *     Id int
 *     Importance int
 *     Subordinates []int
 * }
 */

func getImportance(employees []*Employee, id int) int {
	// dfs + map
	var res int
	m := make(map[int]*Employee)
	for _, x := range employees {
		m[x.Id] = x
	}
	var count func(id int)
	count = func(id int) {
		res += m[id].Importance
		for _, y := range m[id].Subordinates {
			count(y)
		}	
	}
	count(id)
	return res
}
// @lc code=end

