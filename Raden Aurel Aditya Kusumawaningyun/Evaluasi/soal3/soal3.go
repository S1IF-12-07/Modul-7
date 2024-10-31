package main

import "fmt"

func main() {
	var kirat int

	fmt.Scanln(&kirat)

	dinar := kirat / (10 * 10 * 6)
	kirat %= (10 * 10 * 6)
	dirham := kirat / (10 * 6)
	kirat %= (10 * 6)
	fals := kirat / 6
	kirat %= 6

	fmt.Printf("%d\n", dinar)
	fmt.Printf("%d\n", dirham)
	fmt.Printf("%d\n", fals)
	fmt.Printf("%d\n", kirat)
}
