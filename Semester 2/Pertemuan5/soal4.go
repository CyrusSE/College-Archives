package main

import "fmt"

func pangkat(x, y float64) float64 {
	var hasil float64
	hasil = 1
	for i := 1.0; i <= y; i++ {
		hasil *= x
	}
	return hasil
}

func main() {
	var x, hasil float64
	fmt.Scan(&x)
	hasil = 6*pangkat(x, 3) + 4*pangkat(x, 2) + 2
	fmt.Println(hasil)
}
