package main

import "fmt"

func main() {
	var pendudukawal, kelahiran, imigrasi, kematian, emigrasi int
	var jumlahpenduduk int
	fmt.Print("Masukan jumlah untuk penduduk awal : ")
	fmt.Scan(&pendudukawal)
	fmt.Print("Masukan jumlah untuk kelahiran : ")
	fmt.Scan(&kelahiran)
	fmt.Print("Masukan jumlah untuk imigrasi : ")
	fmt.Scan(&imigrasi)
	fmt.Print("Masukan jumlah untuk kematian : ")
	fmt.Scan(&kematian)
	fmt.Print("Masukan jumlah untuk emigrasi : ")
	fmt.Scan(&emigrasi)
	jumlahpenduduk = (pendudukawal + ((kelahiran + imigrasi) - (kematian + emigrasi)))
	fmt.Print(jumlahpenduduk)
}
