package main

import (
	"fmt"
)

func main() {
	var x, y int
	
	fmt.Print("Masukkan bilangan positif x: ")
	fmt.Scan(&x)
	
	fmt.Print("Masukkan bilangan positif y: ")
	fmt.Scan(&y)

    hasil:= 0
	for i := x; i <= y; i++ {

	    hasil += i
	}
	fmt.Print(hasil)

}