package main

import "fmt"

func main() {
	var p, l int
	var luas, keliling int
	fmt.Scan(&p, &l)
	luas = p * l
	keliling = (2 * p) + (2 * l)
	fmt.Println(keliling, luas)
}
