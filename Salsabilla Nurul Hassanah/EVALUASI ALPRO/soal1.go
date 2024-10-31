//NAMA: SALSABILLA NURUL HASSANAH
//KELAS: S1IF-12-07
//NIM: 103112430256

package main

import (
	"fmt"
)

func main() {
	var x, y int

	fmt.Scan(&x)
	fmt.Scan(&y)

	n := y - x + 1
	sum := n * (x + y) / 2

	fmt.Println(sum)
}