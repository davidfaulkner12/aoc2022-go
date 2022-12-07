package problems

import ()

func day06(s string, l int) int64 {
	i := int64(0)

Window:
	for i = 0; i < int64(len(s)-l); i++ {
		window := s[i : i+int64(l)]
		for j := 0; j < l; j++ {
			n := window[j]
			for k := j + 1; k < l; k++ {
				if n == window[k] {
					continue Window
				}
			}
		}
		break
	}
	return i + int64(l)
}

func day06Prob1(s string) interface{} {
	return day06(s, 4)
}

func day06Prob2(s string) interface{} {
	return day06(s, 14)
}

var Day06 = ProblemSet{
	Problem1: day06Prob1,
	Problem2: day06Prob2,
}
