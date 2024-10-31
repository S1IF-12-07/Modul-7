package main
import "fmt"
func main(){
	var n, keluaran int
	fmt.Scan(&n)
	for i:=1; i<=n; i++{
		keluaran = i*i
		fmt.Println(keluaran)
	}
}