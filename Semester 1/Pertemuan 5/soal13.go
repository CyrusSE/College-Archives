package main

import "fmt"

func main() {
	var n int
	var huruf string
	fmt.Print("Masukan n : ")
	fmt.Scan(&n)
	fmt.Print("Masukan huruf : ")
	fmt.Scan(&huruf)
	for i := 1; i < n; i++ {
		fmt.Println(huruf)
	}
}
