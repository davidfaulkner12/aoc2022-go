package problems

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
)

type Stack[T any] struct {
	Vals []T
}

func (s *Stack[T]) Push(val T) {
	s.Vals = append(s.Vals, val)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.Vals) == 0 {
		var zero T
		return zero, false
	}
	top := s.Vals[len(s.Vals)-1]
	s.Vals = s.Vals[:len(s.Vals)-1]
	return top, true
}

func parseStack(ss []string) []Stack[byte] {
	// Heuristic but haven't validated the math
	lenStacks := len(ss[0])/4 + 1

	stacks := make([]Stack[byte], lenStacks)

	// We want the actual bottom row
	for i := len(ss) - 2; i >= 0; i-- {
		row := ss[i]
		for j := 1; j <= len(row); j += 4 {
			if row[j] != ' ' {
				stacks[(j-1)/4].Push(row[j])
			}
		}
	}

	return stacks
}

type Move struct {
	Quantity int64
	Source   int64
	Target   int64
}

func parseMove(s string) (Move, error) {
	// We're going to cheat and just extract all the numbers
	re := regexp.MustCompile("[0-9]+")

	ss := re.FindAllString(s, -1)

	if len(ss) != 3 {
		return Move{}, errors.New("Did not find three numbers")
	}

	xs := make([]int64, 3)

	for i, s := range ss {
		x, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return Move{}, err
		}
		xs[i] = x
	}

	return Move{xs[0], xs[1], xs[2]}, nil
}

func parseMoves(ss []string) []Move {
	ms := make([]Move, 0, len(ss))
	for _, s := range ss {
		m, err := parseMove(s)
		if err != nil {
			continue
		}
		ms = append(ms, m)
	}
	return ms
}

func day05Prob1(s string) interface{} {
	parts := strings.Split(s, "\n\n")

	stackString := strings.Split(parts[0], "\n")
	moveString := strings.Split(parts[1], "\n")

	stacks := parseStack(stackString)
	moves := parseMoves(moveString)

	for _, m := range moves {
		for i := int64(0); i < m.Quantity; i++ {
			x, _ := stacks[m.Source-1].Pop()
			stacks[m.Target-1].Push(x)
		}
	}

	result := make([]byte, len(stacks))

	for i, s := range stacks {
		result[i], _ = s.Pop()
	}

	return string(result)

}

func day05Prob2(s string) interface{} {
	parts := strings.Split(s, "\n\n")

	stackString := strings.Split(parts[0], "\n")
	moveString := strings.Split(parts[1], "\n")

	stacks := parseStack(stackString)
	moves := parseMoves(moveString)

	for _, m := range moves {
		is := Stack[byte]{}
		for i := int64(0); i < m.Quantity; i++ {
			x, _ := stacks[m.Source-1].Pop()
			is.Push(x)
		}
		for i := int64(0); i < m.Quantity; i++ {
			x, _ := is.Pop()
			stacks[m.Target-1].Push(x)
		}
	}

	result := make([]byte, len(stacks))

	for i, s := range stacks {
		result[i], _ = s.Pop()
	}

	return string(result)

}

var Day05 = ProblemSet{
	Problem1: day05Prob1,
	Problem2: day05Prob2,
}
