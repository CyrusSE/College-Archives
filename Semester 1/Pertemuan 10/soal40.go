package main

import "fmt"

func main() {
	var a, b, c, d, e float64
	fmt.Print("Masukkan 5 bilangan rill : ")
	fmt.Scan(&a, &b, &c, &d, &e)
	if a < b && b < c && c < d && d < e {
		fmt.Print("Stabil naik")
	} else if a > b && b > c && c > d && d > e {
		fmt.Print("Stabil turun")
	} else {
		fmt.Print("Tidak stabil")
	}
}
