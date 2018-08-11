package parity

func Parity(num int) int {
	parity := 0

	for num > 0 {
		parity ^= 1
		num &= num - 1
	}

	return parity
}
