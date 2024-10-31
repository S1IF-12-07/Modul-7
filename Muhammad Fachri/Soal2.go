package main

import "fmt"

func main() {
    var n int
    fmt.Println("Masukkan nilai n ")
    fmt.Scanln(&n)
    if n <= 0 {
        fmt.Println("Nilai n bilangan bulat positif")
        return
    }
    for i := 1; i <= n; i++ {
        fmt.Printf("%d ", i*i)
    }
	fmt.Println()
}