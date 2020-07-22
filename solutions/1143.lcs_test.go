package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestCommonSubsequence1143(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, LongestCommonSubsequence1143("abcde", "ace"))
	assert.Equal(3, LongestCommonSubsequence1143("abc", "abc"))
	assert.Equal(0, LongestCommonSubsequence1143("", ""))
	assert.Equal(0, LongestCommonSubsequence1143("", "abc"))
}
