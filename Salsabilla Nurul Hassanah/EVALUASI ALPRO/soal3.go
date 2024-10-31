//NAMA: SALSABILLA NURUL HASSANAH
//KELAS: S1IF-12-07
//NIM: 103112430256

package main

import "fmt"

func main() {
	var qirat int

	fmt.Scan(&qirat)

	dinar := qirat / 600           
	dirham := (qirat % 600) / 60    
	fals := (qirat % 60) / 6        
	qiratSisa := qirat % 6          

	fmt.Println(dinar, dirham, fals, qiratSisa)
}