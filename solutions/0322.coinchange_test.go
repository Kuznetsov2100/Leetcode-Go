package solutions

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoinChange0322(t *testing.T) {
	assert := assert.New(t)

	assert.Equal(3, CoinChange0322([]int{1, 2, 5}, 11))
	assert.Equal(-1, CoinChange0322([]int{2}, 3))
}
