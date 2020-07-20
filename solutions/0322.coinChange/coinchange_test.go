package problem0322

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, CoinChange([]int{1, 2, 5}, 11))
	assert.Equal(-1, CoinChange([]int{2}, 3))
}
