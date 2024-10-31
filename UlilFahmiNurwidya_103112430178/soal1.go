package main

import (
	"fmt"
)

func jumlahkan(x, y int) int {
	total := 0
	for i := x; i <= y; i++ {
		total += i
	}
	return total
}

func main() {
	var x, y int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)

	if x <= y {
		hasil := jumlahkan(x, y)
		fmt.Printf("Hasil penjumlahan dari %d sampai %d adalah: %d\n", x, y, hasil)
	} else {
		fmt.Println("Error: x harus lebih kecil atau sama dengan y.")
	}
}
