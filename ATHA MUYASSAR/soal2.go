package main
import "fmt"

func main()  {
	var n int
	fmt.Print("masukan n : ")
	fmt.Scan(&n)

	for i := 1; i<=n; i++{
		fmt.Printf("%d=%d\n", i, i*i)
	}
}