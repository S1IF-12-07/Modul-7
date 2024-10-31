package main

import "fmt"

func main() {
	var N int

	fmt.Print("Masukkan nilai N: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Nilai N harus positif.")
		return
	}

	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", i*i)
	}

	fmt.Println()
}
