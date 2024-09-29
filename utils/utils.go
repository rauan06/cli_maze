package utils

func IsValidInput(n rune) bool {
	digits := []rune{'0', '1', '2', '3'}

	for _, digit := range digits {
		if n == digit {
			return true
		}
	}

	return false
}
