package main

import "fmt"

func main() {
	var qirat int
	fmt.Println("Masukkan nilai qirat :")
	fmt.Scanln(&qirat)
	fmt.Println("Hasil :")
	fmt.Scanln("Hasil")

	dinar := qirat / 600
	dirham := (qirat % 600) / 60
	fals := (qirat % 60) / 6
	qiratSisa := qirat % 6

	fmt.Println(dinar, dirham, fals, qiratSisa)
}
