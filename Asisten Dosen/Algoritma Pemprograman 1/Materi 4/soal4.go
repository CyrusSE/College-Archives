package main

import "fmt"

type karyawan struct {
	nama      string
	gajiPokok float64
	tunjangan float64
	potongan  float64
}

func main() {
	var data karyawan
	var TotalGaji float64
	fmt.Println("Informasi Karyawan:")
	fmt.Print("Nama: ")
	fmt.Scan(&data.nama)
	fmt.Print("Gaji Pokok: ")
	fmt.Scan(&data.gajiPokok)
	fmt.Print("Tunjangan: ")
	fmt.Scan(&data.tunjangan)
	fmt.Print("Potongan: ")
	fmt.Scan(&data.potongan)
	TotalGaji = data.gajiPokok + data.tunjangan - data.potongan
	fmt.Printf("Total Gaji: Rp %.2f", TotalGaji)
}
