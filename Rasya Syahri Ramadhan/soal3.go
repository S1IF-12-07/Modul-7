package main

import "fmt"
func konversi(qirat int) (int, int, int, int) {
	const qiratPerFals = 6
	const falsPerDirham = 10
	const dirhamPerDinar = 10
	const qiratPerDinar = qiratPerFals * falsPerDirham * dirhamPerDinar
	dinar := qirat / qiratPerDinar
	sisaQirat := qirat % qiratPerDinar

	dirham := sisaQirat / (qiratPerFals * falsPerDirham)
	sisaQirat %= (qiratPerFals * falsPerDirham)

	fals := sisaQirat / qiratPerFals
	sisaQirat %= qiratPerFals

	hasilqirat := sisaQirat

	return dinar, dirham, fals, hasilqirat
}
func main() {
	var qiratt int
	fmt.Print("Masukkan jumlah qirat: ")
	fmt.Scan(&qiratt)
	dinar, dirham, fals, qirat := konversi(qiratt)
	fmt.Printf("%d qirat: Dinar = %d Dirham = %d Fals = %d Qirat = %d\n", qiratt, dinar, dirham, fals, qirat)
}
