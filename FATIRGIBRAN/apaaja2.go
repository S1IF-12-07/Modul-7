package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Printf("%d ", i*i)
	}
	fmt.Println()
}
