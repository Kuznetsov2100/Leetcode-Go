/*
 * @lc app=leetcode.cn id=327 lang=golang
 *
 * [327] 区间和的个数
 */

// @lc code=start
func countRangeSum(nums []int, lower int, upper int) int {
	var res int
	n := len(nums)
	// 计算前缀和 preSum，以及后面统计时会用到的所有数字 allNums
    allNums := make([]int, 1, 3*n+1)
    preSum := make([]int, n+1)
    for i, v := range nums {
        preSum[i+1] = preSum[i] + v
        allNums = append(allNums, preSum[i+1], preSum[i+1]-lower, preSum[i+1]-upper)
    }

    // 将 allNums 离散化
    sort.Ints(allNums)
    k := 1
    kth := map[int]int{allNums[0]: k}
    for i := 1; i <= 3*n; i++ {
        if allNums[i] != allNums[i-1] {
            k++
            kth[allNums[i]] = k
        }
	}
	
	root := buildSegmentTree(1, k)
	update(kth[0], root)
	for _, sum := range preSum[1:] {
		left, right := kth[sum-upper], kth[sum-lower]
        res += query(left, right, root)
        update(kth[sum], root)
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

