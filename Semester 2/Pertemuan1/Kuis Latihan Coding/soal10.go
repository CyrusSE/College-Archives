package main

import "fmt"

func main() {
	var x int
	var angka1, angka2, hasil int
	fmt.Scan(&x)
	angka1 = x / 10
	angka2 = x % 10
	hasil = (angka1 * 1000) + (angka1 * 100) + (angka2 * 10) + angka2
	fmt.Println(hasil)
}
