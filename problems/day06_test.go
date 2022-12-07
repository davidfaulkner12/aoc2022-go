package problems

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var day06_example = []string{
	"mjqjpqmgbljsphdztnvjfqwrcgsmlb",
	"bvwbjplbgvbhsrlpgdmjqwftvncz",
	"nppdvjthqldpwncqszvftbrmjlhg",
	"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
	"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
}

func TestDay06P1(t *testing.T) {
	wants := []int64{
		7,
		5,
		6,
		10,
		11,
	}

	for i, s := range day06_example {
		want := wants[i]
		assert.Equal(t, want, day06Prob1(s))
	}
}

func TestDay06P2(t *testing.T) {
	wants := []int64{
		19,
		23,
		23,
		29,
		26,
	}

	for i, s := range day06_example {
		want := wants[i]
		assert.Equal(t, want, day06Prob2(s))
	}
}

func TestDay06ActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/day06.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day06.Problem1(s)
	assert.Equal(t, int64(1640), ans)

	ans = Day06.Problem2(s)
	assert.Equal(t, int64(3613), ans)
}
