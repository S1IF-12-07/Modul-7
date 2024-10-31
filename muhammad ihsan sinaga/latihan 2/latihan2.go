package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scanln(&n)

	fmt.Print("Keluaran: ")
	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println()
}
