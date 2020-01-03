package main

import (
	"fmt"

	"time"

	"github.com/itzmeanjan/number-theory/numbertheory"
)

func main() {
	time.Now()
	init := time.Now()
	fmt.Printf("Triangular Number ( formula ) : %d\nTime taken : %v\n", numbertheory.GetTriangularNumber(100), time.Now().Sub(init))
	init = time.Now()
	fmt.Printf("Triangular Number ( iterative ) : %d\nTime taken : %v\n", numbertheory.GetTriangularNumberAlt(100), time.Now().Sub(init))
	init = time.Now()
	fmt.Printf("Triangular Number ( recursive ) : %d\nTime taken : %v\n", numbertheory.GetTriangularNumberRec(100), time.Now().Sub(init))
}
