package main

import "fmt"

func main() {
	var n int
	fmt.Println("Masukan nilai n")
	fmt.Scanln(&n)
	for i := 1; i <= n; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}
