package tools_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	. "github.com/davidfaulkner12/aoc2022/tools"
)

func TestRead(t *testing.T) {
	type test struct {
		input string
		want  []int64
	}

	tests := []test{
		{input: "38", want: []int64{38}},
		{input: "abc def", want: []int64{}},
		{input: "", want: []int64{}},
		{input: "6294967297\n7\n8\nbob\n323", want: []int64{6294967297, 7, 8, 323}},
	}

	for _, tc := range tests {
		got, ok := ReadNums(tc.input)
		assert.Nil(t, ok, "Error was not nil")
		assert.Equal(t, tc.want, got, "Slices were not equal")
	}
}
