package main

import "fmt"

func main() {
  var qirat int
  var sisa int
  fmt.Println("Masukkan jumlah qirat ")
  fmt.Scanln(&qirat)
  dinar := qirat / 600
  sisa = qirat % 600
  dirham := sisa / 60 
  sisa = sisa % 60
  fals := sisa / 6 
  sisa = sisa % 6
  fmt.Printf("%d %d %d %d\n", dinar, dirham, fals, sisa)
}