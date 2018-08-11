package parity_test

import (
	"testing"
	"github.com/disiqueira/playground/parity-bit"
)

func TestParity(t *testing.T) {
	testCases := map[int]int {
		11 : 1,
		2345: 1,
		1313123 : 0,
		3456787654: 0,
		98765432345: 1,
	}

	for num, expected := range testCases {
		res := parity.Parity(num)
		if res != expected {
			t.Errorf("parity fail for %d, %d got, %d expected", num, res, expected)
		}
	}
}
