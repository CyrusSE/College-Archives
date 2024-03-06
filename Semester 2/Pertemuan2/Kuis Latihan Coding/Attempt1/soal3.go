package main

import "fmt"

func faktor(x, y int) bool {
	return y%x == 0
}

func main() {
	var b1, b2 int
	var jenis bool
	fmt.Scan(&b1, &b2)
	jenis = faktor(b1, b2)
	fmt.Println(jenis)
}
