package main

import "fmt"

func gunung(x, y, n int) int {
	jumlah_letusan := 0
	sisa_tahun := n
	i := 1
	for sisa_tahun >= x {
		if i%2 != 0 {
			sisa_tahun -= x
			jumlah_letusan++
		} else {
			sisa_tahun -= y
			jumlah_letusan++
		}
		i++
	}
	return jumlah_letusan
}

func main() {
	var x, y, n int
	fmt.Scan(&x, &y, &n)
	fmt.Println(gunung(x, y, n))
}
