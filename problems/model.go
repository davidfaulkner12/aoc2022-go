package problems

type Problem func(x string) interface{}

type ProblemSet struct {
	Problem1 Problem
	Problem2 Problem
}
