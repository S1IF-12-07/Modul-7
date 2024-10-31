package main

import "fmt"

func main() {
	var x, y int
	var hasil int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scanln(&x)
    fmt.Print("Masukkan nilai y: ")
    fmt.Scanln(&y)

	for i := x; i <= y; i++{
		hasil += i
	}
	fmt.Print("Hasilnya adalah: ", hasil)
    
}