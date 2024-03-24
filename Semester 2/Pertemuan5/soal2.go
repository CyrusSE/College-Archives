package main

import "fmt"

func penjumlahan(x int) int {
	if x < 10 {
		x += 1
	} else {
		x += 2
	}
	return x
}

func tambah(x *int) {
	if *x < 10 {
		*x += 1
	} else {
		*x += 2
	}

}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(penjumlahan(n))
	tambah(&n)
	fmt.Println(n)
}
