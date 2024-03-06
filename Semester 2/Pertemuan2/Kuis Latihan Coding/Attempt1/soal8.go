package main

import "fmt"

func nilaiAkhir(x, y float64) float64 {
	return (x + y) / 2
}

func main() {
	var tulis, tugas, nakhir float64
	fmt.Scan(&tulis, &tugas)
	nakhir = nilaiAkhir(tulis, tugas)
	fmt.Printf("%.2f", nakhir)
}
