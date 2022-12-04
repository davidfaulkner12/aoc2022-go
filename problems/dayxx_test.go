package problems

import (
	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
	// "os"
	"testing"
)

var dayX_example = `2-4,6-8
2-3,4-5
5-7,7-9
2-8,3-7
6-6,4-6
2-6,4-8
`

func TestDayXP1(t *testing.T) {
	ans := dayXProb1(dayX_example)
	assert.Equal(t, int64(0), ans)
}

func TestDayXP2(t *testing.T) {
	ans := dayXProb2(dayX_example)
	assert.Equal(t, int64(0), ans)
}

// func TestDayXActualProblem(t *testing.T) {
// b, err := os.ReadFile("../data/dayX.txt")
// require.Nil(t, err)
// s := string(b)

// ans := DayX.Problem1(s)
// assert.Equal(t, int64(0), ans)

// ans = DayX.Problem2(s)
// assert.Equal(t, int64(0), ans)
// }
