package problems

import (
	"github.com/alecthomas/participle/v2"
)

// Heavily indebted to https://github.com/alecthomas/participle/blob/master/_examples/expr/main.go

// Parser

type Rochambeau int64

const (
	Paper    Rochambeau = 2
	Rock     Rochambeau = 1
	Scissors Rochambeau = 3
)

var rochambeauMap = map[string]Rochambeau{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

func (r *Rochambeau) Capture(s []string) error {
	*r = rochambeauMap[s[0]]
	return nil
}

type RochambeauGame struct {
	Left  Rochambeau `@("A" | "B" | "C" | "X" | "Y" | "Z")`
	Right Rochambeau `@("A" | "B" | "C" | "X" | "Y" | "Z")`
}

type RochambeauGames struct {
	Games []*RochambeauGame `@@*`
}

var parser = participle.MustBuild[RochambeauGames]()

// Eval
func (r RochambeauGame) Score() (int64, int64) {
	var win, draw int64 = 6, 3

	var left, right int64

	switch {
	case r.Left == r.Right:
		left = draw
		right = draw
	case r.Left.defeats() == r.Right:
		left = win
	default:
		right = win
	}

	return left + int64(r.Left), right + int64(r.Right)
}

func (r Rochambeau) defeats() Rochambeau {
	switch r {
	case Paper:
		return Rock
	case Rock:
		return Scissors
	case Scissors:
		return Paper
	default:
		panic("Unexpected rochambeau value!")
	}
}

func (r Rochambeau) losesTo() Rochambeau {
	switch r {
	case Paper:
		return Scissors
	case Rock:
		return Paper
	case Scissors:
		return Rock
	default:
		panic("Unexpected rochambeau value!")
	}
}

func (r RochambeauGame) ScorePart2() (int64, int64) {
	var win, draw int64 = 6, 3

	// Rock (X) is a loss
	// Paper (Y) is a draw
	// Scissors (Z) is a win
	switch r.Right {
	case Rock:
		return win + int64(r.Left), int64(r.Left.defeats())
	case Paper:
		return draw + int64(r.Left), draw + int64(r.Left)
	case Scissors:
		return int64(r.Left), win + int64(r.Left.losesTo())
	default:
		panic("Unexpected rochambeau value!")
	}
}

// Solution

func runProblem(s string, scorePart2 bool) int64 {
	games, err := parser.ParseString("", s)
	if err != nil {
		panic(err)
	}

	var rightTotal int64

	for _, g := range games.Games {
		var right int64
		if scorePart2 {
			_, right = g.ScorePart2()
		} else {
			_, right = g.Score()
		}
		rightTotal += right
	}

	return rightTotal
}

func day2Prob1(s string) int64 {
	return runProblem(s, false)
}

func day2Prob2(s string) int64 {
	return runProblem(s, true)
}

var Day02 = ProblemSet{
	Problem1: day2Prob1,
	Problem2: day2Prob2,
}
