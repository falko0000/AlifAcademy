package commission

func Calculate(amount int) int {

	minAmount := 500000
	if amount <= minAmount {
		return 0
	}

	result := amount / 20000
	return result
}
