package main

import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	fmt.Println("Nilai x harus lebih kecil atau sama dengan y.")

	var hasil int

	for i := x; i <= y; i++ {
		hasil += i
	}

	fmt.Printf("Hasil penjumlahan dari %d sampai %d adalah: %d\n", x, y, hasil)
}
