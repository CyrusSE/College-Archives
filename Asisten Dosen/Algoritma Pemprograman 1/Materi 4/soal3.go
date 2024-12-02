package main

import "fmt"

type badan struct {
	nama   string
	berat  float64
	tinggi float64
}

func main() {
	var data badan
	var BMI float64
	fmt.Println("Informasi BMI:")
	fmt.Print("Nama: ")
	fmt.Scan(&data.nama)
	fmt.Print("Berat: ")
	fmt.Scan(&data.berat)
	fmt.Print("Tinggi: ")
	fmt.Scan(&data.tinggi)
	BMI = data.berat / (data.tinggi * data.tinggi)
	fmt.Printf("BMI: %.2f", BMI)
}
