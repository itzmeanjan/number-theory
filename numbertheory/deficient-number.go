package numbertheory

// IsDeficientNumber - Fetches all proper factors of a number & finds sum of them
// equating that sum with number, gets us whether its Deficient, Abundant or Perfect Number
//
// if sum is greater than actual number, it's abundant number
//
// if sum is lesser than actual number, it's deficient number
//
// if sum is equal to actual number, it's perfect number
func IsDeficientNumber(num int) bool {
	factors := Factorize(num)
	sum := 0
	for i := 0; i < len(factors); i++ {
		sum += factors[i]
	}
	if sum < num {
		return true
	}
	return false
}
