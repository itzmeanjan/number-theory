package main

import (
	"fmt"

	"github.com/itzmeanjan/number-theory/numbertheory"
)

func main() {
	/*
		time.Now()
		init := time.Now()
		fmt.Printf("\n100th Triangular Number ( formula ) : %d\tdone in %v\n", numbertheory.GetTriangularNumber(100), time.Now().Sub(init))
		init = time.Now()
		fmt.Printf("100th Triangular Number ( iterative ) : %d\tdone in %v\n", numbertheory.GetTriangularNumberAlt(100), time.Now().Sub(init))
		init = time.Now()
		fmt.Printf("100th Triangular Number ( recursive ) : %d\tdone in %v\n\n", numbertheory.GetTriangularNumberRec(100), time.Now().Sub(init))
		if numbertheory.IsSquareNumber(81) {
			fmt.Println("81 is a SQUARE NUMBER")
		} else {
			fmt.Println("81 isn't a SQUARE NUMBER")
		}
		if numbertheory.IsSquareNumber(82) {
			fmt.Println("82 is a SQUARE NUMBER")
		} else {
			fmt.Println("82 isn't a SQUARE NUMBER")
		}
		fmt.Printf("\nSum of reciprocal of first 10 Triangular Numbers : %f\n", numbertheory.GetSumOfReciprocalsOfFirstNTriangularNumbers(10))
		fmt.Printf("Sum of reciprocal of first 100 Triangular Numbers : %f\n", numbertheory.GetSumOfReciprocalsOfFirstNTriangularNumbers(100))
		fmt.Printf("Sum of reciprocal of first 1000 Triangular Numbers : %f\n\n", numbertheory.GetSumOfReciprocalsOfFirstNTriangularNumbers(1000))
		fmt.Printf("First 10 Triangular Square Numbers : %v\n", numbertheory.GetFirstXTriangularSquareNumbers(10))
		fmt.Printf("Is sum of 780 & 990 TRIANGULAR = %v\nIs diff of 780 & 990 TRIANGULAR = %v\n", numbertheory.IsTriangularNumber(780+990), numbertheory.IsTriangularNumber(990-780))
	*/
	fmt.Println("\nTriangular Numbers which are sum of two other triangular numbers ( from first 1000 ) :")
	for _, i := range numbertheory.GetTriangularNumberWhichAreSumOfTriangularNumbers(1000) {
		fmt.Printf("\t%d -- sum(%d, %d)\n", i.TriNum, i.TriNumOne, i.TriNumTwo)
	}
}
