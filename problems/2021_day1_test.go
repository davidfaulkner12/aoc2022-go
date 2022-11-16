package problems

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var example = `199
200
208
210
200
207
240
269
260
263
`

func TestP1(t *testing.T) {
	ans := p1(example)
	assert.Equal(t, int64(7), ans)
}

func TestP2(t *testing.T) {
	ans := p2(example)
	assert.Equal(t, int64(5), ans)
}

func TestActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/2021-day1.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day2021_01.Problem1(s)
	assert.Equal(t, int64(1759), ans)

	ans = Day2021_01.Problem2(s)
	assert.Equal(t, int64(1805), ans)
}
