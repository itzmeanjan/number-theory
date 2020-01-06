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

// IsSumOfTwoTriangularNumbers - Checks whether `num` can be represented
//  as sum of two other triangular numbers or not
//
// for ease of calculation we need to first generate all triangular numbers
// which are lesser than `n`, if `num` is `n-th` triangular number,
// and slice which is holding those triangular numbers, need to be passed as
// argument
//
// another argument to be taken, position of `num` in triangular number series
// i.e. 6 is 3rd triangular number in series of all triangular numbers
//
// In return, it'll get us a 3-element tuple, where first element denotes
// whether check is successful or not, and next two values are those two triangular numbers
// when summed forms this triangular number
func IsSumOfTwoTriangularNumbers(num int, idx int, triNumArr []int) (bool, int, int) {
	success := false
	numOne := 0
	numTwo := 0
	for i := 0; i < idx-1; i++ {
		for j := i + 1; j < idx-1; j++ {
			tmp := triNumArr[i] + triNumArr[j]
			if tmp > num {
				break
			}
			if tmp == num && IsTriangularNumber(tmp) {
				success = true
				numOne = triNumArr[i]
				numTwo = triNumArr[j]
				break
			}
		}
		if success {
			break
		}
	}
	return success, numOne, numTwo
}

// SpecialTriangularNumber - This exportable struct holds
// a triangular number in its first field,
// and two other different triangular numbers,
// such that when summed up, will get us first triangular number
// i.e. first triangular number can be represented  as a sum of
// other two triangular numbers
type SpecialTriangularNumber struct {
	TriNum    int
	TriNumOne int
	TriNumTwo int
}

// GetTriangularNumberWhichAreSumOfTriangularNumbers - Returns a slice holding
// all triangular numbers ( from first `n` triangular numbers ), which can
// be represented as sum of two other triangular numbers
func GetTriangularNumberWhichAreSumOfTriangularNumbers(n int) []SpecialTriangularNumber {
	triNumArr := GenerateFirstNTriangularNumbers(n)
	filteredTriNumArr := []SpecialTriangularNumber{}
	for i := 0; i < cap(triNumArr); i++ {
		if flag, v1, v2 := IsSumOfTwoTriangularNumbers(triNumArr[i], i+1, triNumArr); flag {
			filteredTriNumArr = append(filteredTriNumArr, SpecialTriangularNumber{triNumArr[i], v1, v2})
		}
	}
	return filteredTriNumArr
}

// TriangularNumberPair - ...
type TriangularNumberPair struct {
	One int
	Two int
}

// findPossibleTriNumComb - Finds out possible triangular number pair
// which satisfies following condition :
//
// If a & b are two triangular numbers such that a != b & b > a
// then b + a & b - a, need to be triangular too
//
// that what this function tries to find out from a slice of triangular numbers
func findPossibleTriNumComb(triNumArr []int) []TriangularNumberPair {
	filteredTriNumPair := []TriangularNumberPair{}
	for i := 0; i < len(triNumArr); i++ {
		for j := i + 1; j < len(triNumArr); j++ {
			sum := triNumArr[j] + triNumArr[i]
			diff := triNumArr[j] - triNumArr[i]
			if IsTriangularNumber(sum) && IsTriangularNumber(diff) {
				filteredTriNumPair = append(filteredTriNumPair, TriangularNumberPair{triNumArr[i], triNumArr[j]})
			}
		}
	}
	return filteredTriNumPair
}

// putNonDuplicates - Given two slices, we'll return only a
// single slice ( first one as per this implementation )
// in which all values of second slice will be appended, if and only if
// that certain value is not already present in first slice
func putNonDuplicates(mainArr []TriangularNumberPair, auxArr []TriangularNumberPair, fillInit int) int {
	checkUpto := fillInit
	for i := 0; i < len(auxArr); i++ {
		found := false
		for j := 0; j < checkUpto; j++ {
			if auxArr[i] == mainArr[j] {
				found = true
				break
			}
		}
		if !found {
			mainArr[fillInit] = auxArr[i]
			fillInit++
		}
	}
	return fillInit
}

// GetXTriangularNumbersWhichAreSumAndDiffOfTwoOtherTriangularNumbers - ...
func GetXTriangularNumbersWhichAreSumAndDiffOfTwoOtherTriangularNumbers(x int) []TriangularNumberPair {
	triNumPairArr := make([]TriangularNumberPair, x)
	generatedTriNumArr := []int{}
	curTriNumIdx := 1
	foundTriNumPairPos := 0
	for foundTriNumPairPos < cap(triNumPairArr) {
		generatedTriNumArr = append(generatedTriNumArr, GetTriangularNumber(curTriNumIdx))
		curTriNumIdx++
		foundTriNumPairPos = putNonDuplicates(triNumPairArr, findPossibleTriNumComb(generatedTriNumArr), foundTriNumPairPos)
	}
	return triNumPairArr
}
