package main

import "fmt"

func main() {
	var x, y int
	var hasil int
	fmt.Println("Masukan nilai x")
	fmt.Scanln(&x)
	fmt.Println("Masukan nilai y")
	fmt.Scanln(&y)

	for i := x; i <= y; i++ {
		hasil += i
	}
	fmt.Printf("%d", hasil)
}
