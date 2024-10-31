package main

import "fmt"

func main() {
	var N int
	fmt.Print("Masukan bilangan bulat positif N: ")
	fmt.Scanln(&N)

	fmt.Print("Keluaran: ")
	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println()
}