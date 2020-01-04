package numbertheory

// IsSquareNumber - Checks whether number passed as parameter is
// a square number or not
//
// what we do for checking is, we'll sum up first `n` odd numbers,
// & if that sum is equal to our target number, then it's of course a Square Number.
//
// but if that sum crosses target number, we'll simply get out of that loop
func IsSquareNumber(num int) bool {
	sum := 0
	success := false
	for i := 1; i < num; i += 2 {
		sum += i
		if sum == num {
			success = true
			break
		} else if sum > num {
			break
		}
	}
	return success
}
