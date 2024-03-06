package main

import "fmt"

func kelilingPersegiPanjang(x, y int) int {
	return 2 * (x + y)
}

func main() {
	var panjang, lebar int
	fmt.Scan(&panjang, &lebar)
	fmt.Println(kelilingPersegiPanjang(panjang, lebar))
}
