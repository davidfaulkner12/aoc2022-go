package problems

type Problem func(x string) int64

type ProblemSet struct {
	Problem1 Problem
	Problem2 Problem
}
