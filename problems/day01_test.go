package problems

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var day1_example = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestDay1P1(t *testing.T) {
	ans := day1Prob1(day1_example)
	assert.Equal(t, int64(24000), ans)
}

func TestDayP2(t *testing.T) {
	ans := day1Prob2(day1_example)
	assert.Equal(t, int64(45000), ans)
}

func TestDay1ActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/day1.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day01.Problem1(s)
	assert.Equal(t, int64(70720), ans)

	ans = Day01.Problem2(s)
	assert.Equal(t, int64(207148), ans)
}
