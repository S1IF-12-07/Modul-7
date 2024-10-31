package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	sum := 0
	for i := x; i <= y; i++ {
		sum += i
	}

	fmt.Println(sum)
}
