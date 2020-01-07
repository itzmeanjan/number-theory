package numbertheory

// GetOblongNumber - Returns N-th oblong number,
// which is sum of first N even numbers
func GetOblongNumber(n int) int {
	return n * (n + 1)
}

// IsOblongNumber - Checks whether a given number is oblong number
// or not, by summing up all even numbers ( starting from 2 )
// until `sum` equals to / crosses `num` ( whichever happens first ),
// breaks out of loop
//
// Finally makes a simple check by equating value of `sum` & `num`,
// if they are equal to each other, then it's an Oblong Number
func IsOblongNumber(num int) bool {
	sum := 0
	for i := 2; i < num; i += 2 {
		sum += i
		if sum == num || sum > num {
			break
		}
	}
	return sum == num
}
