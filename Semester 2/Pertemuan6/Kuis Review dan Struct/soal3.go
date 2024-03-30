package main

import "fmt"

func belanja(x, y int) int {
	/* Fungsi mengembalikan nilai dari total harga yang harus dibayar */
	if y < 0 {
		return 0
	} else if x == 1 {
		return y * 2980
	} else if x == 2 {
		return y * 4500
	} else if x == 3 {
		return y * 9980
	} else if x == 4 {
		return y * 4490
	} else if x == 5 {
		return y * 6870
	} else {
		return 0
	}
}

func main() {
	var noProduk, banyakBelanja, total int
	fmt.Scan(&noProduk, &banyakBelanja)
	total = belanja(noProduk, banyakBelanja)
	fmt.Println(total)
}
