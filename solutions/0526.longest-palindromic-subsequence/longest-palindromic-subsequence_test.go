package problem0516

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubseq(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, LongestPalindromeSubseq("bbbab"))
	assert.Equal(2, LongestPalindromeSubseq("cbbd"))

}
