package numbertheory

import "math"

// IsPrime - Checks whether a given number is prime or not
// If not prime, then it's a composite number
func IsPrime(num int) bool {
	check := true
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			check = false
			break
		}
	}
	return check
}

// GeneratePrimesUnderX - Generates all primes numbers under `X` ( starting from 2 )
func GeneratePrimesUnderX(x int) []int {
	primes := []int{2}
	for i := 3; i < x; i++ {
		check := true
		sqrt := int(math.Sqrt(float64(i)))
		for j := 0; j < len(primes) && primes[j] <= sqrt; j++ {
			tmp := primes[j]
			if i%tmp == 0 {
				check = false
				break
			}
		}
		if check {
			primes = append(primes, i)
		}
	}
	return primes
}
