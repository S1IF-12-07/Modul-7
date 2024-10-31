package main
import "fmt"

func main() {
	var x, y int

	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Print("Masukkan nilai y: ")
	fmt.Scan(&y)

	fmt.Println("Nilai x harus lebih kecil atau sama dengan y.")

	var hasil int

	for i := x; i <= y; i++ {
		hasil += i
	}

	fmt.Println("Hasil penjumlahan dari" , x , "sampai", y ,"adalah", hasil )
}