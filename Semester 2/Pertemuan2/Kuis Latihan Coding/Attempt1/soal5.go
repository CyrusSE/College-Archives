package main

import "fmt"

func kelilinglingkaran(x float64) float64 {
	return 2 * 3.14 * x
}

func main() {
	var radius, keliling float64
	fmt.Scan(&radius)
	keliling = kelilinglingkaran(radius)
	fmt.Println(keliling)
}
