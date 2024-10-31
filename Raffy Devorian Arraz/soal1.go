package main

import "fmt"

func main() {
    var x, y int
    fmt.Print("Masukkan nilai x: ")
    fmt.Scanln(&x)
    fmt.Print("Masukkan nilai y: ")
    fmt.Scanln(&y)
    if x > y {
        fmt.Println("Nilai x harus lebih kecil atau sama dengan y")
        return
    }
    jumlah := (y - x + 1) * (x + y) / 2
    fmt.Printf("Jumlah bilangan dari %d sampai %d adalah %d\n", x, y, jumlah)
}