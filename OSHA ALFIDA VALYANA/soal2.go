package main

import "fmt"

func main() {
	var N int
	fmt.Println("Masukkan bilangan bulat positif N :")
	fmt.Scanln(&N)
	fmt.Println("Hasil :")
	fmt.Scanln("Hasil")

	for i := 1; i <= N; i++ {
		fmt.Print(i*i, " ")
	}
}
