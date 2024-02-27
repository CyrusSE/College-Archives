package main

import "fmt"

func main() {
	var x, penjumlahan float64
	fmt.Print("Masukan angka untuk x : ")
	fmt.Scan(&x)
	penjumlahan = ((x * x * x) + 3*x) / ((x * x * x * x) - 3*(x*x) + 4)
	fmt.Println("Total Penjumlahan : ", penjumlahan)
	fmt.Println("x = ", x)
}
