# Advent of Code (aoc)

For 2022, I'm doing Advent of Code in Go!

## Usage

### Build

Build the single AOC command:

```bash
$ go build ./cmd/aoc
```

### Run

```bash
# "Usage: aoc problem file"
$ ./aoc 01 data/day1.txt
Problem 1: <answer>
Problem 2: <answer>
```

### Test

I use Go testing and the [testify library](https://github.com/stretchr/testify).

```bash
$ go test ./...
?       github.com/davidfaulkner12/aoc2022/cmd/aoc  [no test files]
ok      github.com/davidfaulkner12/aoc2022/problems (cached)
ok      github.com/davidfaulkner12/aoc2022/tools    (cached)
```

## Design Notes

### Add a Problem Set

There are two steps necessary to add a problem set. Firstly, add two Go files, `<yyyy>-day<day>.go` and the equivalent test. Create a data file in `data/` for your problem input.

Secondly, add an entry to the `Problems` map in `cmd/aoc/main.go`. By convention, problems for 2022 are given a key of "DD", eg., "01". Problems for other years are given by "YYYY-DD", eg., "2021-01".

### Testing and Imports

For problems, we do NOT use a separate `<package>_test`, but instead put the tests into the package. Given the variety of problems, we have many situations where it is convenient and important to test the internal API, as opposed to the very general `string -> int64` public interface.

We use `<package>_test` and the [`.` import pattern](https://github.com/golang/go/wiki/CodeReviewComments#import-dot) in non-problem tests.

## License

Copyright © 2022 David Wang-Faulkner

This program and the accompanying materials are made available under the
terms of the MIT License. See [LICENSE](LICENSE).

