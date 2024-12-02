package main

import "fmt"

type investasi struct {
	HargaBeli   float64
	HargaJual   float64
	JumlahSaham float64
}

func main() {
	var data investasi
	var investasiawal, totalpenjualan, keuntungankotor, biayatransaksi, pajakkeuntungan, keuntunganbersih float64

	fmt.Println("Informasi Investasi Saham:")
	fmt.Print("Harga Beli: Rp ")
	fmt.Scan(&data.HargaBeli)
	fmt.Print("Harga Jual: Rp ")
	fmt.Scan(&data.HargaJual)
	fmt.Print("Jumlah Saham: ")
	fmt.Scan(&data.JumlahSaham)

	investasiawal = data.HargaBeli * data.JumlahSaham
	fmt.Printf("Total Investasi Awal: Rp %.2f\n", investasiawal)

	totalpenjualan = data.HargaJual * data.JumlahSaham
	fmt.Printf("Total Penjualan: Rp %.2f\n", totalpenjualan)

	keuntungankotor = totalpenjualan - investasiawal
	fmt.Printf("Keuntungan Kotor: Rp %.2f\n", keuntungankotor)

	biayatransaksi = 0.002 * totalpenjualan
	fmt.Printf("Biaya Transaksi: Rp %.2f\n", biayatransaksi)

	if keuntungankotor < 0 {
		pajakkeuntungan = 0
	} else {
		pajakkeuntungan = 0.1 * keuntungankotor
	}
	fmt.Printf("Pajak Keuntungan %.2f\n", pajakkeuntungan)

	keuntunganbersih = keuntungankotor - biayatransaksi - pajakkeuntungan
	fmt.Printf("Keuntungan Bersih: Rp %.2f", keuntunganbersih)

}
