package numbertheory

// GetTriangularNumber - returns triangular number by position ( in series )
func GetTriangularNumber(pos int) int {
	return pos * (pos + 1) / 2
}

// GetTriangularNumberAlt - Alternative method for getting triangular number by position
// works by iteratively summing up integers upto requested position of triangular number
func GetTriangularNumberAlt(pos int) int {
	tmp := 0
	for i := 1; i < (pos + 1); i++ {
		tmp += i
	}
	return tmp
}

// GetTriangularNumberRec - Calculates triangular number at a certain position
// in recursive fashion
func GetTriangularNumberRec(pos int) int {
	if pos == 1 {
		return 1
	}
	return pos + GetTriangularNumberRec(pos-1)
}

// GenerateFirstNTriangularNumbers - Generates first `n` triangular numbers
// using prewritten functions, and returns them as an array
func GenerateFirstNTriangularNumbers(n int) []int {
	buffer := make([]int, n)
	for i := 0; i < cap(buffer); i++ {
		buffer[i] = GetTriangularNumber(i + 1)
	}
	return buffer
}

// GetSumOfReciprocalsOfFirstNTriangularNumbers - Returns sum of reciprocal
// of first `n` triangular numbers, which tends to `2`, with increase in `n`
func GetSumOfReciprocalsOfFirstNTriangularNumbers(n int) float64 {
	buffer := GenerateFirstNTriangularNumbers(n)
	sum := 0.0
	for i := 0; i < cap(buffer); i++ {
		sum += float64(1) / float64(buffer[i])
	}
	return sum
}

// IsTriangularNumber - Checks whether a number is triangular or not
// i.e. sums up all integers upto some value `n`, and if that sum is
// equal to this number, then its a triangular number
//
// while continuing this summing up, if sum crosses `number`,
// then we'll simply break out of that loop
func IsTriangularNumber(num int) bool {
	if num == 1 {
		return true
	}
	sum := 0
	success := false
	for i := 1; i < num; i++ {
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

// GetFirstXTriangularSquareNumbers - Returns an slice holding
// first `X` triangular numbers which are square numbers too
func GetFirstXTriangularSquareNumbers(x int) []int {
	buffer := make([]int, x)
	triIndex := 1
	i := 0
	for i < x {
		tmp := GetTriangularNumber(triIndex)
		triIndex++
		if IsSquareNumber(tmp) {
			buffer[i] = tmp
			i++
		}
	}
	return buffer
}
