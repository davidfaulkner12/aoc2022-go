package problems

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	//"fmt"
	"strings"
	"testing"
)

var day05_example = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2 `

func TestDay05StackParser(t *testing.T) {
	parts := strings.Split(day05_example, "\n\n")

	stackString := strings.Split(parts[0], "\n")

	stacks := parseStack(stackString)

	want := []Stack[byte]{
		Stack[byte]{[]byte("ZN")},
		Stack[byte]{[]byte("MCD")},
		Stack[byte]{[]byte("P")},
	}

	assert.Equal(t, want, stacks)

}

func TestDay05MoveParser(t *testing.T) {
	parts := strings.Split(day05_example, "\n\n")

	moveString := strings.Split(parts[1], "\n")

	moves := parseMoves(moveString)

	want := []Move{
		Move{1, 2, 1},
		Move{3, 1, 3},
		Move{2, 2, 1},
		Move{1, 1, 2},
	}

	assert.Equal(t, want, moves)

}

func TestDay05P1(t *testing.T) {
	ans := day05Prob1(day05_example)
	assert.Equal(t, "CMZ", ans)
}

func TestDay05P2(t *testing.T) {
	ans := day05Prob2(day05_example)
	assert.Equal(t, "MCD", ans)
}

func TestDay05ActualProblem(t *testing.T) {
	b, err := os.ReadFile("../data/day05.txt")
	require.Nil(t, err)
	s := string(b)

	ans := Day05.Problem1(s)
	assert.Equal(t, "PTWLTDSJV", ans)

	ans = Day05.Problem2(s)
	assert.Equal(t, "WZMFVGGZP", ans)
}
