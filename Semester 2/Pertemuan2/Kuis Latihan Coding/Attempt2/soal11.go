package main

import "fmt"

func jumlahGenap(a, b int) int {
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
		if i%2 == 0 {
			hasil += i
		}
	}
	return hasil
}

func main() {

	var bil1, bil2 int
	fmt.Scan(&bil1, &bil2)
	fmt.Println(jumlahGenap(bil1, bil2))
}
