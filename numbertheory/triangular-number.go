package numbertheory

// GetTriangularNumber - returns triangular number by position ( in series )
func GetTriangularNumber(pos int) int64 {
	return int64(pos) * int64(pos+1) / 2
}

// GetTriangularNumberAlt - Alternative method for getting triangular number by position
// works by iteratively summing up integers upto requested position of triangular number
func GetTriangularNumberAlt(pos int) int64 {
	var tmp int64 = 0
	for i := 1; i < (pos + 1); i++ {
		tmp += int64(i)
	}
	return tmp
}

// GetTriangularNumberRec - Calculates triangular number at a certain position
// in recursive fashion
func GetTriangularNumberRec(pos int) int64 {
	if pos == 1 {
		return int64(1)
	}
	return int64(pos) + GetTriangularNumberRec(pos-1)
}
