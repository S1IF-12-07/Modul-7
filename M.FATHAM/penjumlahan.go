package main

import "fmt"

func main() {
	var x, y int
	var hasil int
	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		hasil += i
		fmt.Printf("%d ", hasil)
	}
	fmt.Printf("%d ", hasil)
}
