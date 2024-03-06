package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var hasil int
	fmt.Scan(&a, &b, &c, &d, &e)
	hasil = a + b + c - d - e
	fmt.Print(hasil)
}
