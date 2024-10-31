package main

import "fmt"

func main() {
	var x, y int

	fmt.Scanln(&x, &y)

	penjumlahan := 0
	for i := x; i <= y; i++ {
		penjumlahan += i
	}

	fmt.Println(penjumlahan)
}
