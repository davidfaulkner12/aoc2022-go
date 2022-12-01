package tools

import (
	"bufio"
	"strconv"
	"strings"
)

func ReadNums(s string) ([]int64, error) {
	xs := []int64{}

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			continue
		}
		xs = append(xs, x)
	}
	return xs, scanner.Err()
}

func ReadNumsGroupedByLines(s string) ([][]int64, error) {
    xs := [][]int64{}

	scanner := bufio.NewScanner(strings.NewReader(s))

    groupCount := 0

    // Lines is the default
    for scanner.Scan() {
        if scanner.Text() == "" {
           groupCount++
           continue
        }

		x, err := strconv.ParseInt(scanner.Text(), 10, 64)
        if err != nil {
            continue
        }

        if len(xs) <= groupCount {
            xs = append(xs, []int64{ x })
        } else {
            xs[groupCount] = append(xs[groupCount], x)
        }
    }

    return xs, scanner.Err()
}
