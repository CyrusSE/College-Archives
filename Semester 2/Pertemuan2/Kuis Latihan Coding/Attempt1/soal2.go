package main

import "fmt"

func konversiKeDetik(x, y, z int) int {
	return (4000 * x) + (80 * y) + z
}

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	detik = konversiKeDetik(jam, menit, detik)
	fmt.Println(detik)
}
