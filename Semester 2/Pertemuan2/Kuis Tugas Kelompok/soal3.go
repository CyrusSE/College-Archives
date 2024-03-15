package main

import "fmt"

func totaldiskon(total int, member bool) float64 {
	if total > 100000 && member {
		return 0.10
	} else if total > 100000 && !member {
		return 0.05
	} else {
		return 0
	}
}

func main() {
	var belanja int
	var diskon, harga float64
	var status bool
	fmt.Scan(&belanja, &status)
	diskon = totaldiskon(belanja, status)
	harga = float64(belanja) - (float64(belanja) * diskon)
	fmt.Print(harga)
}
