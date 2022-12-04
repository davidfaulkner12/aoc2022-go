package problems

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var day4_example = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestSplitDoubleRange(t *testing.T) {
	r, err := splitRange("2-4,6-8")
	assert.Nil(t, err)
	assert.Equal(t, DualRange{Range{2, 4}, Range{6, 8}}, r)
}

func TestDay4P1(t *testing.T) {
	ans := day4Prob1(day4_example)
	assert.Equal(t, int64(2), ans)
}

func TestDay4P2(t *testing.T) {
	ans := day4Prob2(day4_example)
	assert.Equal(t, int64(4), ans)
}

func TestDay4ActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/day4.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day04.Problem1(s)
	assert.Equal(t, int64(305), ans)

	ans = Day04.Problem2(s)
	assert.Equal(t, int64(811), ans)
}
