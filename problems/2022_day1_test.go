package problems

import (
	"github.com/stretchr/testify/assert"
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
	ans := p1("bob")
	assert.Equal(t, int64(103943), ans, "p1 result was unexpected")
}
