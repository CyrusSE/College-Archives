package main

import "fmt"

func hargaSetelahDiskon(a, b int, check bool) float64 {
	if check == true && a >= 100000 {
		diskon := float64(a) * (float64(b) / 100)
		return float64(a) - diskon
	} else {
		return float64(a)
	}
}

func main() {
	var harga, diskon int
	var member bool
	var hargaAkhir float64
	fmt.Scan(&harga, &diskon, &member)
	hargaAkhir = hargaSetelahDiskon(harga, diskon, member)
	fmt.Println(hargaAkhir)
}
