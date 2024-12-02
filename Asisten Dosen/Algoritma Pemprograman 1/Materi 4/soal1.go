package main

import "fmt"

type Transaksi struct {
	namaBarang  string
	jumlah      int
	hargaSatuan float64
}

func main() {
	var data Transaksi
	var totalHarga float64
	fmt.Println("Informasi Transaksi:")
	fmt.Print("Nama Barang: ")
	fmt.Scan(&data.namaBarang)
	fmt.Print("Jumlah: ")
	fmt.Scan(&data.jumlah)
	fmt.Print("Harga Satuan: Rp")
	fmt.Scan(&data.hargaSatuan)
	totalHarga = float64(data.jumlah) * data.hargaSatuan
	fmt.Printf("Total Harga: Rp %.2f", totalHarga)
}
