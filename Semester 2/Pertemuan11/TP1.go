package main

import "fmt"

// Deklarasi konstanta NMAX dengan nilai 10
const NMAX int = 10

// Deklarasi tipe alias tabInt untuk array bilangan bulat dengan
// kapasitas maksimum NMAX
type tabInt [NMAX]int

func main() {

	// Deklarasi variabel
	var data tabInt
	var nData int
	var x1 int

	// Baca bilangan yang dicari
	fmt.Scan(&x1)

	// Baca data array
	fmt.Scan(&nData)
	bacaData(&data, &nData)

	// Cetak data
	cetakData(data, nData)

	// Pencarian bilangan tertentu dengan sequential search dan frekuensinya
	fmt.Print("Hasil pencarian: ")
	if sequentialSearch(data, nData, x1) {
		fmt.Printf("Bilangan ditemukan. Terdapat %d bilangan %d.\n", frekuensiBilangan(data, nData, x1), x1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	/*
		IS: Array bilangan bulat A dan n terdefinisi sembarang
		FS: Array bilangan bulat A terisi data sebanyak n elemen.
			n tidak boleh lebih dari NMAX.
	*/
	if *n > NMAX {
		*n = NMAX
	}
	for i := 0; i < *n; i++ {
		fmt.Scan(&A[i])
	}
}

func cetakData(A tabInt, n int) {
	/*
		IS: Array bilangan A dan banyak elemen n terdefinisi
		FS: Tercetak di layar bilangan dengan format:
			"Data Bilangan: <e1> <e2> <e3> ..<en>"
	*/
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func frekuensiBilangan(A tabInt, n, x int) int {
	/*
		Diberikan array bilangan A sebanyak n elemen dan bilangan x yang
		ingin dicari tahu frekuensinya, fungsi mengembalikan frekuensi
		x dalam array A.
	*/
	var angka int
	for i := 0; i < n; i++ {
		if A[i] == x {
			angka++
		}
	}
	return angka
}

func sequentialSearch(A tabInt, n int, x int) bool {
	/*
		Diberikan array bilangan A dengan banyak elemen n dan bilangan
		yang dicari (x), fungsi mengembalikan boolean true jika ditemukan,
		atau boolean false jika tidak ditemukan.
		Gunakan algoritma pencarian beruntun atau sequential search.
	*/
	var ada bool
	for i := 0; i < n; i++ {
		if A[i] == x {
			ada = true
		}
	}
	return ada
}
