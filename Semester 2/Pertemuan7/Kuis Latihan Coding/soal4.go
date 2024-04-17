package main

import "fmt"

const NMax int = 10000

type tabInt [NMax]int

func InputArray(Tab *tabInt, n int) {
	/*I.S. Terdefinisi bilangan bulat n*/
	/*Proses : Menerima masukan bilangan bulat. Proses input akan berhenti
	  ketika sudah mencapai kapasitas maksimum yaitu n*/
	/*F.S. Array Tab berisi data yang diberikan*/
	var i int
	for i = 0; i < n; i++ {
		fmt.Scan(&Tab[i])
	}
}

func jumlahArrGenap(Tab tabInt, n int) int {
	/*Mengembalikan hasil dari penjumlahan isi di dalam array
	  dengan tipe data integer*/
	var i, total int
	for i = 0; i < n; i++ {
		//jika array bernilai genap maka jumlahkan
		if Tab[i]%2 == 0 {
			total += Tab[i]
		}
	}
	return total
}

func main() {
	var Tab tabInt
	var n int
	fmt.Scan(&n)
	InputArray(&Tab, n)
	fmt.Println(jumlahArrGenap(Tab, n))
}
