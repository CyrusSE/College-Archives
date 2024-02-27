package main

import "fmt"

func main() {
	var x, y, penjumlahan float64
	fmt.Print("Masukan angka untuk x : ")
	fmt.Scan(&x)
	fmt.Print("Masukan angka untuk y : ")
	fmt.Scan(&y)
	penjumlahan = 1.0/(3.0*(x*x)+10.0) + 10.0*y + 7
	fmt.Println("Total Penjumlahan : ", penjumlahan)
	fmt.Println("x = ", x, "\ny = ", y)
}
