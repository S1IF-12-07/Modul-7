package main

import "fmt"

func main() {
	var q int
	fmt.Scan(&q)
	dinar := q / 600
	dirham := (q % 600) / 60
	fals := (q % 60) / 6
	qirat := q % 6
	fmt.Print(dinar, dirham, fals, qirat)
}
