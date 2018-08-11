package hindex

import (
	"math"
	"sort"
)

func HIndex(papers []int) int {
	sort.Ints(papers)

	last := math.MinInt32
	for index, num := range papers {
		if len(papers)-index < num {
			last = num
			break
		}
	}

	return last
}
