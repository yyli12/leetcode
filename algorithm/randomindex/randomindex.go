package randomindex

import (
	"fmt"
	"math/rand"
)

type numRange struct {
	min int
	max int
}

type Solution struct {
	numRanges []*numRange
	maxRange  int
}

func Constructor(weights []int) Solution {
	solution := Solution{
		numRanges: make([]*numRange, len(weights)),
	}
	accumulatedWeight := 0
	for i, weight := range weights {
		solution.numRanges[i] = &numRange{
			min: accumulatedWeight,
		}
		accumulatedWeight += weight
		solution.numRanges[i].max = accumulatedWeight
	}
	solution.maxRange = accumulatedWeight
	return solution
}

func (s *Solution) PickIndex() int {
	pick := rand.Intn(s.maxRange)

	// find pick in which index
	l, r := 0, len(s.numRanges)
	m := (l + r) >> 1
	for {
		numRange := s.numRanges[m]
		fmt.Println(pick, m, numRange)
		if numRange.min <= pick && pick < numRange.max {
			return m
		}
		if numRange.min > pick {
			r = m
		} else {
			l = m + 1
		}
		m = (l + r) >> 1
	}
}
