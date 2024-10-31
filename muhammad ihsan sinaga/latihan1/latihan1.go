package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan x: ")
	fmt.Scanln(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scanln(&y)

	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}

	fmt.Println("Hasil penjumlahan:", sum)
}
