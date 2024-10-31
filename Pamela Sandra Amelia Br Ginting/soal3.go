package main

import "fmt"

func main() {
	var qirat int

	fmt.Print("Masukkan jumlah qirat: ")
	fmt.Scan(&qirat)
	dinar := qirat / (10 * 10 * 6)
	sisa := qirat % (10 * 10 * 6)
	dirham := sisa / (10 * 6)
	sisa = sisa % (10 * 6)
	fals := sisa / 6
	sisa = sisa % 6

	fmt.Printf("Hasil penukaran:\n")
	fmt.Printf("Dinar: %d\n", dinar)
	fmt.Printf("Dirham: %d\n", dirham)
	fmt.Printf("Fals: %d\n", fals)
	fmt.Printf("Qirat: %d\n", sisa)
}
