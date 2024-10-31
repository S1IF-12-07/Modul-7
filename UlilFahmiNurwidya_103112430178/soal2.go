package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Print("Masukkan bilangan bulat positif N: ")
	fmt.Scan(&N)

	if N > 0 {
		for i := 1; i <= N; i++ {
			kuadrat := i * i
			fmt.Println(kuadrat)
		}
	} else {
		fmt.Println("Error: N harus bilangan bulat positif.")
	}
}
