/*
 * @lc app=leetcode.cn id=493 lang=golang
 *
 * [493] 翻转对
 */

// @lc code=start
func reversePairs(nums []int) int {
	// 线段树
	n := len(nums)
	if n <= 1 {
		return 0
	}
	
	var res int
	var allNums []int
	for _, num := range nums {
		allNums = append(allNums, num, 2*num+1)
	}
	sort.Ints(allNums)
	m := make(map[int]int)
	index := 0
	for _, num := range allNums {
		if _, ok := m[num]; !ok {
			m[num] = index
			index++
		}
	}
	root := buildSegmentTree(0, index-1)
	for _, num := range nums {
		res += query(m[2*num+1],index-1, root)
		update(m[num], root)
	}
	return res	
}

type TreeNode struct {
	start, end, count int
	Left,Right *TreeNode
}

func buildSegmentTree(start, end int) *TreeNode {
	if start > end {
		return nil
	}
	mid := start + (end-start)/2
	node := &TreeNode{start:start, end: end}
	if start < end {
		node.Left = buildSegmentTree(start, mid)
		node.Right = buildSegmentTree(mid+1, end)
	}
	return node
}

func update(index int, node *TreeNode) {
	if node == nil {
		return
	}
	if index > node.end || index < node.start {
		return
	}
	node.count++
	update(index, node.Left)
	update(index, node.Right)
}

func query(min, max int, node *TreeNode) int {
	if node == nil || min > max {
		return 0
	}
	if node.start == min && node.end == max {
		return node.count
	}
	mid := node.start + (node.end-node.start)/2
	var ans int
	if max <= mid {
		ans = query(min, max, node.Left)
	} else {
		if min <= mid {
			ans = query(min, mid, node.Left) + query(mid+1, max, node.Right)
		} else {
			ans = query(min, max, node.Right)
		}
	}
	return ans
}
// @lc code=end

