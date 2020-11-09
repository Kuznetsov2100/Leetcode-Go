/*
 * @lc app=leetcode.cn id=307 lang=golang
 *
 * [307] 区域和检索 - 数组可修改
 */

// @lc code=start
// 使用线段树，query和update操作的时间复杂度为O(logn)
type NumArray struct {
	arr []int
	tree []int
	arrLength int
}


func Constructor(nums []int) NumArray {
	n := len(nums)
	numarray := NumArray{
		arr : nums,
		tree : make([]int, 4*n),
		arrLength: n,
	}
	if n > 0 {
		numarray.buildTree(0, 0, n-1)
	}
	return numarray
}


func (this *NumArray) Update(i int, val int)  {
	if this.arrLength > 0 {
		this.updateTree(0, 0, this.arrLength-1, i, val)
	}
}


func (this *NumArray) SumRange(i int, j int) int {
	if i > j || this.arrLength == 0 {
		return 0
	}
	return this.query(0, 0, this.arrLength-1, i, j)
}


func (this *NumArray) buildTree(node, start, end int) {
	if start == end {
		this.tree[node] = this.arr[start]
	} else {
		mid := start + (end-start)/2
		leftChild := 2*node + 1
		rightChild := 2*node + 2
		this.buildTree(leftChild, start, mid)
		this.buildTree(rightChild, mid+1, end)
		this.tree[node] = this.tree[leftChild] + this.tree[rightChild]
	}
}


func (this *NumArray) updateTree(node, start, end, index, value int) {
	if start == end {
		this.arr[index] = value
		this.tree[node] = value
	} else {
		mid := start + (end-start)/2
		leftChild := 2 * node + 1
		rightChild := 2 * node + 2
		if start <= index && index <= mid {
			this.updateTree(leftChild, start, mid, index, value)
		} else if mid < index && index <= end {
			this.updateTree(rightChild, mid+1, end, index, value)
		}
		this.tree[node] = this.tree[leftChild] + this.tree[rightChild]
	}
}


func (this *NumArray) query(node, start, end, L, R int) int {
	if R < start || L > end {
		return 0
	} 
	if L <= start && end <= R {
		return this.tree[node]
	}
	if start == end {
		return this.arr[start]
	}
	mid := start + (end-start)/2
	leftChild := 2 * node + 1
	rightChild := 2 * node + 2
	return this.query(leftChild, start, mid, L, R) + 
		   this.query(rightChild, mid+1, end, L, R)
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
// @lc code=end

