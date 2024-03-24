package main

import "fmt"

func bentuk(x, y int) string {
	if x == y {
		return "Persegi"
	} else {
		return "Persegi Panjang"
	}
}

func luas(x, y int) int {
	return x * y
}

func main() {
	var a, b, hasil int
	var jenis string
	fmt.Scan(&a, &b)
	jenis = bentuk(a, b)
	hasil = luas(a, b)
	fmt.Println("Luas", jenis, hasil)
}
