package main

import "fmt"

const NMax int = 10

type tabInt [NMax]int

func main() {
	var data tabInt
	var nData int
	baca(&data, &nData)
	cetak(data, nData)
	fmt.Println(jumlah(data, nData), rata_rata(data, nData))
}

func baca(A *tabInt, n *int) {
	/*
		IS: Array A dan n terdefinisi sembarang
		Proses: Setiap bilangan yang dibaca ditampung dalam sebuah variabel.
		    Jika bilangan negatif, maka ubah menjadi bilangan positif dan
		    masukan ke dalam array. Pembacaan berakhir jika terbaca bilangan 0.
		FS: Array A sebanyak n elemen berisi bilangan positif
	*/
	var angka int
	for {
		fmt.Scan(&angka)
		if angka == 0 {
			break
		}
		if *n < NMax {
			A[*n] = angka
			*n += 1
		}
	}
}

func cetak(A tabInt, n int) {
	/*
		IS: Array A terdefinisi sebanya kn elemen
		FS: Array A tercetak di layar
	*/
	for i := 0; i < n; i++ {
		if A[i] > 0 {
			fmt.Print(A[i], " ")
		}
	}
	fmt.Println()
}

func jumlah(A tabInt, n int) int {
	/* Mengembalikan jumlah dari nilai bilangan pada elemen array A */
	var total, num int
	for i := 0; i < n; i++ {
		num = A[i]
		if num < 0 {
			num = -num
		}
		total += num
	}
	return total
}

func rata_rata(A tabInt, n int) float64 {
	/* Mengembalikan rata-rata bilangan */
	return float64(jumlah(A, n)) / float64(n)
}
