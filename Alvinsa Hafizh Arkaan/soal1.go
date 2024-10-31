package main

import (
	"fmt"
)

func sumRange(x, y int) int {
	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}
	return sum
}

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("Nilai x harus lebih kecil atau sama dengan y.")
	} else {
		result := sumRange(x, y)
		fmt.Printf("Hasil penjumlahan dari %d sampai %d adalah: %d\n", x, y, result)
	}
}
