package hindex_test

import (
	"github.com/disiqueira/playground/h-index"
	"testing"
)

func TestParity(t *testing.T) {
	testCases := map[int][]int{
		4: {1, 4, 1, 4, 2, 1, 3, 5, 6},
	}

	for expected, papers := range testCases {
		res := hindex.HIndex(papers)
		if res != expected {
			t.Errorf("Wrong H-Index for the %+v, %d got, %d expected", papers, res, expected)
		}
	}
}
