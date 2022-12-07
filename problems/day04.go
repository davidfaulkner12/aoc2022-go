package problems

import (
	"errors"
	//"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Range struct {
	Start int64
	End   int64
}

type DualRange struct {
	R1 Range
	R2 Range
}

func splitRange(s string) (DualRange, error) {
	// We're going to cheat and just extract all the numbers
	re := regexp.MustCompile("[0-9]+")

	ss := re.FindAllString(s, -1)

	if len(ss) != 4 {
		return DualRange{}, errors.New("Did not find four numbers")
	}

	xs := make([]int64, 4)

	for i, s := range ss {
		x, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return DualRange{}, err
		}
		xs[i] = x
	}

	return DualRange{Range{xs[0], xs[1]}, Range{xs[2], xs[3]}}, nil
}

type compareFn func(left, right Range) bool

func contains(left, right Range) bool {
	return (left.Start <= right.Start && left.End >= right.End) || (right.Start <= left.Start && right.End >= left.End)
}

func max(x, y int64) int64 {
	if x > y {
		return x
	}
	return y
}

func min(x, y int64) int64 {
	if x < y {
		return x
	}
	return y
}

func overlaps(left, right Range) bool {
	return max(left.Start, right.Start) <= min(left.End, right.End)
}

func day4(s string, cfn compareFn) int64 {
	rows := strings.Split(s, "\n")

	rs := make([]DualRange, 0, len(rows))

	for _, row := range rows {
		var err error
		dr, err := splitRange(row)
		if err != nil {
			continue
		}
		rs = append(rs, dr)
	}

	var count int64

	for _, dr := range rs {
		r1, r2 := dr.R1, dr.R2
		if cfn(r1, r2) {
			count++
		}
	}

	return count
}

func day4Prob1(s string) interface{} {
	return day4(s, contains)
}

func day4Prob2(s string) interface{} {
	return day4(s, overlaps)
}

var Day04 = ProblemSet{
	Problem1: day4Prob1,
	Problem2: day4Prob2,
}
