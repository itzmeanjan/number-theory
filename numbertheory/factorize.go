package numbertheory

// Factorize - Returns a slice holding all factors of an integer ( passed as argument ),
// except the number itself
func Factorize(num int) []int {
	factors := []int{}
	for i := 1; i <= num/2; i++ {
		if num%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}
