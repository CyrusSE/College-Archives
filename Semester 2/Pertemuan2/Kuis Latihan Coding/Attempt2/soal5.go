package main

import "fmt"

func pemenang(a3, a2, a1, b3, b2, b1 int) string {
	var hasil string
	var hasila, hasilb int
	hasila = a3*3 + a2*2 + a1
	hasilb = b3*3 + b2*2 + b1
	if hasila == hasilb {
		hasil = "imbang"
	} else if hasila > hasilb {
		hasil = "A"
	} else {
		hasil = "B"
	}
	return hasil
}

func main() {
	var a3, a2, a1, b3, b2, b1 int
	var hasil string
	fmt.Scan(&a3, &a2, &a1, &b3, &b2, &b1)
	hasil = pemenang(a3, a2, a1, b3, b2, b1)
	fmt.Println(hasil)
}
