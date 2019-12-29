package main

import (
	"fmt"
	"time"
)

func getNthTriangularNumberBetter(n int) int64 {
	return int64(n) * (int64(n) + 1) / 2
}

func getNthTriangularNumber(n int) int64 {
	var num int64 = 0
	for i := 1; i <= n; i++ {
		num += int64(i)
	}
	return num
}

func main() {
	start := time.Now()
	fmt.Printf("One :\n\t10000th Triangular Number : %d\n\tTime taken : %v\n", getNthTriangularNumber(10000), time.Now().Sub(start))
	start = time.Now()
	fmt.Printf("Two :\n\t10000th Triangular Number : %d\n\tTime taken : %v\n", getNthTriangularNumberBetter(10000), time.Now().Sub(start))
}
