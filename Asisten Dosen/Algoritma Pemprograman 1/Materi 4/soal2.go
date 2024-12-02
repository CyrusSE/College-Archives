package main

import "fmt"

type PersegiPanjang struct {
	panjang  float64
	lebar    float64
	luas     float64
	keliling float64
}

func main() {
	var data PersegiPanjang
	fmt.Println("Informasi Persegi Panjang:")
	fmt.Print("Panjang: ")
	fmt.Scan(&data.panjang)
	fmt.Print("Lebar: ")
	fmt.Scan(&data.lebar)
	data.luas = data.panjang * data.lebar
	fmt.Printf("Luas: %.1f", data.luas)
	data.keliling = 2 * (data.panjang + data.lebar)
	fmt.Printf("\nKeliling: %.1f", data.keliling)
}
