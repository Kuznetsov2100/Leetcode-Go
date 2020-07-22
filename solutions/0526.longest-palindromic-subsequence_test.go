package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLongestPalindromeSubseq0526(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(4, LongestPalindromeSubseq0526("bbbab"))
	assert.Equal(2, LongestPalindromeSubseq0526("cbbd"))

}
