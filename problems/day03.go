package problems

import (
	"errors"
	mapset "github.com/deckarep/golang-set/v2"
	"strings"
)

func rucksackPriority(r rune) (int64, error) {
	cp := int64(r)
	// Technically we should also error on 91-96 :shrug:
	if cp < 65 || cp > 122 {
		return 0, errors.New("Invalid rucksack value")
	}
	if cp >= 65 && cp < 97 {
		return cp - 65 + 27, nil
	}

	return cp - 96, nil
}

func splitRucksack(s string) (mapset.Set[rune], mapset.Set[rune]) {
	splitPoint := len(s) / 2
	s1 := mapset.NewSet[rune]([]rune(s[:splitPoint])...)
	s2 := mapset.NewSet[rune]([]rune(s[splitPoint:])...)

	return s1, s2
}

func day3Prob1(s string) int64 {
	rows := strings.Split(s, "\n")
	scores := make([]int64, len(rows))

	for ridx, rucksack := range rows {
		s1, s2 := splitRucksack(rucksack)
		overlap, ok := s1.Intersect(s2).Pop()

		if !ok {
			continue
		}

		var err error

		scores[ridx], err = rucksackPriority(overlap)

		if err != nil {
			panic("bad input")
		}

	}

	var sum int64

	for _, x := range scores {
		sum += x
	}

	return sum
}

func day3Prob2(s string) int64 {
	rows := strings.Split(s, "\n")

	scores := make([]int64, len(rows)/3)

	for i := 0; i < len(rows)-2; i += 3 {
		s1 := mapset.NewSet[rune]([]rune(rows[i])...)
		s2 := mapset.NewSet[rune]([]rune(rows[i+1])...)
		s3 := mapset.NewSet[rune]([]rune(rows[i+2])...)

		overlap, ok := s1.Intersect(s2).Intersect(s3).Pop()

		if !ok {
			continue
		}

		var err error

		scores[i/3], err = rucksackPriority(overlap)

		if err != nil {
			panic("bad input")
		}
	}

	var sum int64

	for _, x := range scores {
		sum += x
	}

	return sum
}

var Day03 = ProblemSet{
	Problem1: day3Prob1,
	Problem2: day3Prob2,
}
