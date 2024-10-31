package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Print("Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan y: ")
	fmt.Scan(&y)
	for i := x; i <= y; i++ {
		hasil += i
	}
	fmt.Print("hasilnya adalah: ", hasil)
}
