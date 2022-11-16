package main

import (
	"fmt"
	"github.com/davidfaulkner12/aoc2022/problems"
	"os"
)

var Problems = map[string]problems.ProblemSet{
	"2021-01": problems.Day2021_01,
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: aoc problem file")
		os.Exit(1)
	}

	day, ok := Problems[os.Args[1]]

	if !ok {
		fmt.Printf("No value found for %s\n", os.Args[1])
		os.Exit(1)
	}

	b, err := os.ReadFile(os.Args[2])

	if err != nil {
		fmt.Printf("Unable to read file %s\n", os.Args[2])
		os.Exit(1)
	}

	s := string(b)

	fmt.Printf("Problem 1: %d\n", day.Problem1(s))
	fmt.Printf("Problem 2: %d\n", day.Problem2(s))
}
