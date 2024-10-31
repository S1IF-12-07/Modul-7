//NAMA: SALSABILLA NURUL HASSANAH
//KELAS: S1IF-12-07
//NIM: 103112430256

package main

import "fmt"

func main() {
	var N int

	fmt.Scan(&N)

	for i := 1; i <= N; i++ {
		fmt.Print(i*i, " ")
	}
	fmt.Println()
}