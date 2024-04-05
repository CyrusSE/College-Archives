package main

import "fmt"

const NMax int = 100

type tabUsia [NMax]int

func inputArray(T *tabUsia, n *int) {
	/*I.S. Data usia pasien tersedia dalam piranti masukan
	  Proses: Input akan terus berlangsung sampai bilangan yang diinput tidak
	  valid (<=0)
	  F.S. Sejumlah n data usia pasien masuk ke dalam array T
	  Petunjuk : Gunakan while loop untuk melakukan perulangan
	*/
	var angka int
	for {
		fmt.Scan(&angka)
		if angka <= 0 {
			break
		}
		T[*n] = angka
		*n += 1
	}
}

func hitungRata(T tabUsia, n int) float64 {
	/*Mengembalikan rata-rata usia pasien pada array T*/
	var total float64
	for i := 0; i < n; i++ {
		total += float64(T[i])
	}
	return total / float64(n)
}

func main() {
	var T tabUsia
	var n int
	inputArray(&T, &n)
	fmt.Printf("%.0f", hitungRata(T, n))
}
