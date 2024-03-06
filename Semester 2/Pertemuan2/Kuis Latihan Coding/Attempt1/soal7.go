package main

import "fmt"

func cekHitung(b1 int, b2 int, b3 int, jumlah int) bool {
	return b1+b2+b3 == jumlah
}

func main() {
	var b1, b2, b3, jumlah int
	var cek bool

	fmt.Scan(&b1)
	fmt.Scan(&b2)
	fmt.Scan(&b3)
	fmt.Scan(&jumlah)
	cek = cekHitung(b1, b2, b3, jumlah)
	if cek {
		fmt.Println("benar")
	} else {
		fmt.Println("salah")
	}
}
