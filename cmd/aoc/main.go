package main

import (
	"fmt"
	"github.com/davidfaulkner12/aoc2022/problems"
	"os"
)

var ps = map[string]problems.DayProblems{
	"2022-01": problems.Day2022_01,
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: aoc problem")
		os.Exit(1)
	}

	f, ok := ps[os.Args[1]]

	if !ok {
		fmt.Printf("No value found for %s\n", os.Args[1])
		os.Exit(1)
	}

	day := f()
	fmt.Println(day.Problem1("bob"))
	fmt.Println(day.Problem2("dole"))

}
