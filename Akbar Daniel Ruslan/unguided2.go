package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif: ")
	fmt.Scan(&n)

	fmt.Print("Hasil kuadrat dari 1 sampai ", n, ": ")
	for i := 1; i <= n; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}
