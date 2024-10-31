package main

import "fmt"

func main() {
	var n int


	fmt.Print("Masukkan nilai n: ")
	fmt.Scan(&n)

		if n <= 0 {
		fmt.Println("Nilai n harus positif.")
		return
	}

	for i := 1; i <= n; i++ {
		fmt.Printf("%d ", i*i)
	}
	
}