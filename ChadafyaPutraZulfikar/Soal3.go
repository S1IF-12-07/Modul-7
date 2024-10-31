package main

import "fmt"

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah uang dalam satuan Qirat: ")
	fmt.Scan(&qirat)
	dinar := qirat / 600
	qirat %= 600
	dirham := qirat / 60
	qirat %= 60
	fals := qirat / 6
	qirat %= 6
	fmt.Printf("Dalam Dinar adalah : %d\nDalam Dirham adalah: %d\nDalam Fals adalah: %d\nDalam Qirat adalah : %d\n", dinar, dirham, fals, qirat)
}
