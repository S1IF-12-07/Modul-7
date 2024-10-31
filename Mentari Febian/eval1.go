package main

import "fmt"

func main() {
	var x int
	var y int
	var hasil int
	fmt.Print("Masukan nilai x: ")
	fmt.Scan(&x)
	fmt.Print("Masukan nilai y: ")
	fmt.Scan(&y)

	for i := x; i <= y; i++ {
		hasil += i
	}
	fmt.Printf("%d", hasil)
}
