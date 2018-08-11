package main

import (
	"fmt"
	"math"
	)

func main() {
	papers := []int{1, 4, 1, 4, 2, 1, 3, 5, 6}

	count := make([]int, len(papers))

	for _, val := range papers {
		for i := 1; i <= val; i++ {
			count[i]++
		}
	}

	max := math.MinInt32
	for index, val := range count {
		if val >= index {
			max = index
		}
	}

	fmt.Println(max)
}
