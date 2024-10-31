package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Masukkan suku pertama (x): ")
	fmt.Scan(&x)
	fmt.Print("Masukkan suku terakhir (y): ")
	fmt.Scan(&y)

    hasil:= 0
	for i := x; i <= y; i++ {
	    hasil += i
	}
	fmt.Print(hasil)

}