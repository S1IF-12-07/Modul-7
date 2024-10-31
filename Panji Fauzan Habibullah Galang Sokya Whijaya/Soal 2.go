package main

import "fmt"

func main() {
	var x, hasil int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		hasil = i * i
		fmt.Print(hasil, " ")
	}
}
