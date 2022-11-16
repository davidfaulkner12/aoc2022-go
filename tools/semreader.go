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
