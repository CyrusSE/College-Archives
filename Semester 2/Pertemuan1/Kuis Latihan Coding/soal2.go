package main

import "fmt"

func main() {
	var x, y float64
	var hasil float64
	fmt.Scan(&x, &y)
	hasil = 1/(3*(x*x)+10) + 10*y + 7
	fmt.Print(hasil)
}
