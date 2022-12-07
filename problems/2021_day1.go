package problems

import (
	"github.com/davidfaulkner12/aoc2022/tools"
)

func countIncreasing(xs []int64) int64 {
	var count int64 = 0

	for i := 0; i < len(xs)-1; i++ {
		if xs[i] < xs[i+1] {
			count++
		}
	}

	return count
}

func p1(s string) interface{} {
	xs, _ := tools.ReadNums(s)

	return countIncreasing(xs)
}

func p2(s string) interface{} {
	xs, _ := tools.ReadNums(s)

	windows := []int64{}

	for i := 0; i < len(xs)-2; i++ {
		a, b, c := xs[i], xs[i+1], xs[i+2]
		windows = append(windows, a+b+c)
	}

	return countIncreasing(windows)
}

var Day2021_01 = ProblemSet{
	Problem1: p1,
	Problem2: p2,
}
