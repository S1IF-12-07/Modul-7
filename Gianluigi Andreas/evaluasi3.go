package main

import "fmt"

func main() {
    var qirat int
    fmt.Print("Masukkan jumlah uang dalam satuan qirat: ")
    fmt.Scan(&qirat)

    dinar := qirat / 600
    qirat %= 600

    dirham := qirat / 60
    qirat %= 60

    fals := qirat / 6
    qirat %= 6

    fmt.Printf("Dinar: %d, Dirham: %d, Fals: %d, Qirat: %d\n", dinar, dirham, fals, qirat)
}
