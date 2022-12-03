package problems

import (
	"github.com/alecthomas/participle/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

var day2_example = `A Y
B X
C Z
`

func TestDay2Parser(t *testing.T) {
	type test struct {
		input      string
		want       RochambeauGame
		leftScore  int64
		rightScore int64
	}

	tests := []test{
		{"A X", RochambeauGame{Rock, Rock}, 4, 4},
	}

	var parser = participle.MustBuild[RochambeauGame]()

	for _, tc := range tests {
		got, ok := parser.ParseString("", tc.input)
		assert.Nil(t, ok, "Error was not nil")
		assert.Equal(t, tc.want, *got, "Structs were not equal")
		gotLeft, gotRight := got.Score()
		assert.Equal(t, tc.leftScore, gotLeft, "Left score was unexpected")
		assert.Equal(t, tc.rightScore, gotRight, "Right score was unexpected")
	}
}

func TestDay2P1(t *testing.T) {
	ans := day2Prob1(day2_example)
	assert.Equal(t, int64(15), ans)
}

func TestDay2P2(t *testing.T) {
	ans := day2Prob2(day2_example)
	assert.Equal(t, int64(12), ans)
}

func TestDay2ActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/day2.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day02.Problem1(s)
	assert.Equal(t, int64(17189), ans)

	ans = Day02.Problem2(s)
	assert.Equal(t, int64(13490), ans)
}
