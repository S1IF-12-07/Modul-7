package main

import (
	"fmt"
)

func konversiDuit(jumlahQirat int) (int, int, int, int) {
	const qiratPerFals = 6
	const falsPerDirham = 10
	const dirhamPerDinar = 10

	const qiratPerDinar = qiratPerFals * falsPerDirham * dirhamPerDinar

	dinar := jumlahQirat / qiratPerDinar
	sisaQirat := jumlahQirat % qiratPerDinar

	dirham := sisaQirat / (qiratPerFals * falsPerDirham)
	sisaQirat %= (qiratPerFals * falsPerDirham)

	fals := sisaQirat / qiratPerFals
	sisaQirat %= qiratPerFals

	qiratHasil := sisaQirat

	return dinar, dirham, fals, qiratHasil
}

func main() {
	var jumlahQirat int
	fmt.Print("Masukkan jumlah qirat: ")
	fmt.Scan(&jumlahQirat)

	dinar, dirham, fals, qirat := konversiDuit(jumlahQirat)
	fmt.Printf("%d qirat: Dinar = %d, Dirham = %d, Fals = %d, Qirat = %d\n", jumlahQirat, dinar, dirham, fals, qirat)
}
