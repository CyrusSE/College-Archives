package main

import "fmt"

func main() {
	var x, y float64
	var hasil bool
	fmt.Scan(&x, &y)
	hasil = x*x*x == y
	fmt.Print(hasil)
}
