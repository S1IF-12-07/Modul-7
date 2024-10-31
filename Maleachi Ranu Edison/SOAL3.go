package main

import (
	"fmt"
)

func main() {

	var qirat int

	fmt.Scan(&qirat)

	dinar := qirat / (10 * 10 * 6)
	sisaQirat := qirat % (10 * 10 * 6)

	dirham := sisaQirat / (10 * 6)
	sisaQirat = sisaQirat % (10 * 6)

	fals := sisaQirat / 6
	konversiqirat := sisaQirat % 6

	// Output hasil konversi
	fmt.Printf("%d %d %d %d\n", dinar, dirham, fals, konversiqirat)
}
