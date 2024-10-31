package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	if x > y {
		fmt.Println("Nilai x harus lebih kecil atau sama dengan y.")
		return
	}

	n := y - x + 1
	sum := n * (x + y) / 2

	fmt.Printf("Jumlah dari %d hingga %d adalah: %d\n", x, y, sum)
}