package main
import "fmt"

func main()  {
	var dinar, dirham, fals, qirat int
	fmt.Scan(&qirat)
	dinar = qirat / 600
	dirham = (qirat % 600) / 60
	fals = (qirat % 60) / 6
	qirat = qirat % 6
	fmt.Println("dinar", dinar,"dirham" ,dirham, "fals", fals, "qirat", qirat)
}