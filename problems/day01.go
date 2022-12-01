package problems

import (
	"github.com/davidfaulkner12/aoc2022/tools"
  	"container/heap"
)

func day1Prob1(s string) int64 {
    xss, _ := tools.ReadNumsGroupedByLines(s)

    var max_calories int64 = 0

    for _, xs := range xss {
        var sum int64 = 0
        for _, x := range xs {
            sum += x
        }
        if sum > max_calories {
            max_calories = sum
        }
    }

	return max_calories
}


// An IntHeap is a max-heap of ints.
type IntHeap []int64

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int64))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func day1Prob2(s string) int64 {
    xss, _ := tools.ReadNumsGroupedByLines(s)

    h := &IntHeap{}
    heap.Init(h)

    for _, xs := range xss {
        var sum int64 = 0
        for _, x := range xs {
            sum += x
        }
        heap.Push(h, sum)
    }

	return heap.Pop(h).(int64) + heap.Pop(h).(int64) + heap.Pop(h).(int64)
}

var Day01 = ProblemSet{
	Problem1: day1Prob1,
	Problem2: day1Prob2,
}
