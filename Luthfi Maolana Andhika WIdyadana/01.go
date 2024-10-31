package main

import "fmt"

func main() {
	var x, y, hasil int

	fmt.Print("Masukkan x dan y : ")
	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		hasil += i
	}

	fmt.Printf("%d ", hasil)
}
