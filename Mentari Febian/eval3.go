package main

import "fmt"

func main() {
	var qirat int

	fmt.Println("Masukan qirat")
	fmt.Scan(&qirat)

	dinar := qirat / 600
	dirham := (qirat % 600) / 60
	fals := (qirat % 60) / 6
	qirat_sisa := qirat % 6

	fmt.Println(dinar, dirham, fals, qirat_sisa)
}
