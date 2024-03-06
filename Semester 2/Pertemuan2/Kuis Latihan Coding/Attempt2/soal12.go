package main

import "fmt"

func jumlah(a, b int) int {
	var hasil int
	var terbesar, terkecil int
	if a > b {
		terbesar = a
		terkecil = b
	} else {
		terbesar = b
		terkecil = a
	}
	for i := terkecil; i <= terbesar; i++ {
		hasil += i
	}
	return hasil
}

func main() {
	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlah(bil1, bil2))
}
