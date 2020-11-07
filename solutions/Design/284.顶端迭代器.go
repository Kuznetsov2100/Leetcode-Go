/*
 * @lc app=leetcode.cn id=284 lang=golang
 *
 * [284] 顶端迭代器
 */

// @lc code=start
/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *       
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type PeekingIterator struct {
	iterator *Iterator
	value int // 如需适应所有的类型，改为：value interface{}
	peeked bool
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{iterator:iter}
}

func (this *PeekingIterator) hasNext() bool {
    if this.peeked {
		return true
	}
	return this.iterator.hasNext()
}

func (this *PeekingIterator) next() int {
	if this.peeked {
		this.peeked = false
		return this.value
	}
	return this.iterator.next()
}

func (this *PeekingIterator) peek() int {
	if !this.peeked {
		this.value = this.iterator.next()
		this.peeked = true
	}
	return this.value
}
// @lc code=end

