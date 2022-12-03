package problems

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var day3_example = `vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw
`

func TestDay3Priority(t *testing.T) {
	type test struct {
		input rune
		want  int64
	}

	tests := []test{
		{'A', 27},
		{'a', 1},
		{'Z', 52},
		{'z', 26},
	}

	for _, tc := range tests {
		got, ok := rucksackPriority(tc.input)
		assert.Nil(t, ok, "Error was not nil")
		assert.Equal(t, tc.want, got)
	}
}

func TestSplitRucksack(t *testing.T) {
	s1, s2 := splitRucksack("vJrwpWtwJgWrhcsFMMfFFhFp")
	overlap, _ := s1.Intersect(s2).Pop()
	assert.Equal(t, 'p', overlap)
}

func TestDay3P1(t *testing.T) {
	ans := day3Prob1(day3_example)
	assert.Equal(t, int64(157), ans)
}

func TestDay3P2(t *testing.T) {
	ans := day3Prob2(day3_example)
	assert.Equal(t, int64(70), ans)
}

func TestDay3ActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/day3.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day03.Problem1(s)
	assert.Equal(t, int64(7831), ans)

	ans = Day03.Problem2(s)
	assert.Equal(t, int64(2683), ans)
}
