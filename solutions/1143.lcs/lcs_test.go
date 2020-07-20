package problem1143

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, LongestCommonSubsequence("abcde", "ace"))
	assert.Equal(3, LongestCommonSubsequence("abc", "abc"))
	assert.Equal(0, LongestCommonSubsequence("", ""))
	assert.Equal(0, LongestCommonSubsequence("", "abc"))
}
