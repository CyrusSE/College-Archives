package main

import "fmt"

func main() {
	var n, hasil int
	fmt.Print("Masukkan nilai n : ")
	fmt.Scan(&n)
	hasil = 1
	for i := n; i >= 1; i-- {
		hasil = hasil * i
	}
	fmt.Print(hasil)
}
