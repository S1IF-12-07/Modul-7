package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	fmt.Print("Keluaran: ")
	for i := 1; i <= N; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}
