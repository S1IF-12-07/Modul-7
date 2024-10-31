package main

import (
	"fmt"
)

func main () {
	var n int
	fmt.Scanln(&n)
	for i := 1 ; i <= n; i++ {
		hasil := i * i
		fmt.Println(hasil)
	}
	
}