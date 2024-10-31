package main
import "fmt"

func main()  {
	var x, y, jmlh int
	fmt.Print("masukan x : ")
	fmt.Scan(&x)
	fmt.Print("masukan y : ")
	fmt.Scan(&y)
	for i := x; i<=y; i++{
		jmlh += i
	}
	fmt.Print(jmlh)
}