package main

import "fmt"

func main() {
	var n int
	fmt.Print("Masukkan total angka: ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Printf("%d kuadrat = %d\n", i, i*i)
	}
}
