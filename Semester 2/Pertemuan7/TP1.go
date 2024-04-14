package main

import "fmt"

const NMax int = 10

type tabInt [NMax]int

func main() {
	var nilaiAkhir tabInt
	var banyakNilai int
	bacaNilai(&nilaiAkhir, &banyakNilai)
	cetakNilai(nilaiAkhir, banyakNilai)
}

func bacaNilai(NA *tabInt, n *int) {
	/*
		IS: Nilai akhir (NA) dan banyak elemen (n) terdefinisi sembarang,
		Proses: Membaca banyak elemen (n). Jika n > NMAX, maka n bernilai NMAX.
			Membaca nilai sebanyak n dan memasukkannya ke dalam array nilai akhir (NA).
		FS: Nilai akhir (NA) berisi nilai. Banyak elemen (n) berisi nilai NMAX, jika n > NMAX
	*/
	fmt.Scan(n)
	if *n > NMax {
		*n = NMax
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&NA[i])
	}
}

func cetakNilai(NA tabInt, n int) {
	/*
		IS: Nilai akhir (NA) dan banyak elemen (n) terdefinisi
		FS: Mencetak sebanyak n elemen array nilai akhir (NA)
			dengan format: "Nilai yang terdapat pada array: <n1> <n2> <n3> ..."
			Jika array kosong, maka cetak: "Array kosong"
	*/
	if n == 0 {
		fmt.Println("Array kosong")
	} else {
		fmt.Print("Nilai yang terdapat pada array: ")
		for i := 0; i < n; i++ {
			fmt.Print(NA[i], " ")
		}
	}
}
