package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum0001(t *testing.T) {
	assert := assert.New(t)

	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}
	assert.Equal(expected, TwoSum0001(nums, target))

	nums1 := []int{1, 2, 3}
	target1 := 6
	expected1 := []int(nil)
	assert.Equal(expected1, TwoSum0001(nums1, target1))
}
