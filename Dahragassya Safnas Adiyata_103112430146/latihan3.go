package main

import (
	"fmt"
)

func main() {
	var qirat int
	fmt.Print("Masukkan jumlah satuan qirat: ")
	fmt.Scan(&qirat)

		fals := qirat / 6
		qirat = qirat % 6

		dirham := fals / 10
		fals = fals % 10

		dinar := dirham / 10
		dirham = dirham % 10

	fmt.Print(dinar, dirham, fals, qirat)
}