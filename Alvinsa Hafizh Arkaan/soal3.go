package main

import (
	"fmt"
)

func main() {
	var qirat int
	fmt.Print("Masukkan uang dalam satuan qirat: ")
	fmt.Scan(&qirat)

	// Konversi dari qirat ke satuan yang lebih besar
	fals := qirat / 6
	qirat = qirat % 6

	dirham := fals / 10
	fals = fals % 10

	dinar := dirham / 10
	dirham = dirham % 10

	// Menampilkan hasil konversi
	fmt.Printf("Hasil konversi: \nDinar: %d\nDirham: %d\nFals: %d\nQirat: %d\n", dinar, dirham, fals, qirat)
}
