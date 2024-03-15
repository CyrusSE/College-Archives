package main

import "fmt"

func konversi(j, m, d int) int {
	j = j * 3600
	m = m * 60
	return j + m + d
}

func main() {
	var jam, menit, detik int
	fmt.Scan(&jam, &menit, &detik)
	detik = konversi(jam, menit, detik)
	fmt.Print("Hasil konversi = ", detik, " detik")
}
