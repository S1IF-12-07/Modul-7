package main

import (
	"fmt"
)

func penjumlahan(x, y int) int {
	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}
	return sum
}

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	if x > y {
		fmt.Println("x lebih kecil atau sama dengan y.")
	} else {
		result := penjumlahan(x, y)
		fmt.Printf("%d\n", result)
	}
}
