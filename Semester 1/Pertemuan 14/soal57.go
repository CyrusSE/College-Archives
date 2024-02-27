package main

import "fmt"

func main() {
	var n, lebar int
	fmt.Print("Masukkan banyak daun : ")
	fmt.Scan(&n)
	terbesar := 0
	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan lebar daun ke ", i, " : ")
		fmt.Scan(&lebar)
		if lebar > terbesar {
			terbesar = lebar
		}
	}
	fmt.Print(terbesar)
}
