package main

import "fmt"

func main() {
	var n float32
	var hasil, angka float32
	fmt.Print("Masukan jumlah baris : ")
	fmt.Scan(&n)
	for i := float32(1); i <= n; i++ {
		fmt.Print("Masukkan angka ke ", i, " : ")
		fmt.Scan(&angka)
		hasil += angka
	}
	fmt.Print(hasil / n)
}
