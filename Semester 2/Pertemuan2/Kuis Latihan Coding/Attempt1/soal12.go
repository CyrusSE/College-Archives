package main

import "fmt"

func persentase(penumpang, kapasitasBus int) float64 {
	return float64(penumpang) / float64(kapasitasBus) * 100
}

func valid(persentase float64) bool {
	return persentase >= 50 && persentase <= 75
}

func main() {
	var kapasitasBus, jumlahPenumpang int
	var cek bool
	var persen float64
	fmt.Scan(&kapasitasBus)
	fmt.Scan(&jumlahPenumpang)
	persen = persentase(jumlahPenumpang, kapasitasBus)
	cek = valid(persen)
	if cek {
		fmt.Println("berangkat")
	} else {
		fmt.Println("tidak berangkat")
	}
}
