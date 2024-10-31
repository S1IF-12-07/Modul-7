package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	for i := 1; i <= x; i++ {
		fmt.Printf("%d", i*i)
	}
}
