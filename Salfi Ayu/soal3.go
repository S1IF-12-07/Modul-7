package main

import "fmt"

func main() {
	var Qirat int

	fmt.Print("Masukkan nilai Qirat: ")
	fmt.Scan(&Qirat)

	
	Dinar := Qirat / 600            
	Dirham := (Qirat % 600) / 60    
	Fals := (Qirat % 60) / 6        
	QiratSisa := Qirat % 6          

	
	fmt.Println("Hasil yang didapatkan: ", Dinar, Dirham, Fals, QiratSisa)
}