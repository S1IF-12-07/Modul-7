package main

import "fmt"

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah qirat: ")
	fmt.Scanln(&qirat)

	dinar := qirat / (10 * 10 * 6)
	qirat %= (10 * 10 * 6)
	dirham := qirat / (10 * 6)
	qirat %= (10 * 6)
	fals := qirat / 6
	qirat %= 6

	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", dirham)
	fmt.Printf("Fals: %d\n", fals)
	fmt.Printf("Qirat: %d\n", qirat)
}
