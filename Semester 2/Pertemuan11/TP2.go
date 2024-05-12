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
	var x2 int
	var idx int

	// Baca bilangan yang dicari
	fmt.Scan(&x2)

	// Baca data array
	bacaData(&data, &nData)

	// Cetak data
	cetakData(data, nData)

	// Pencarian bilangan tertentu dengan binary search
	fmt.Print("Hasil pencarian: ")
	binarySearch(data, nData, x2, &idx)

	// Jika idx tidak bernilai -1, maka bilangan yang dicari ditemukan
	if idx != -1 {
		fmt.Printf("Bilangan %d ditemukan pada posisi ke-%d\n", data[idx], idx+1)
	} else {
		fmt.Println("Bilangan tidak ditemukan.")
	}
}

func bacaData(A *tabInt, n *int) {
	/*
	   IS: Array bilangan bulat A dan n terdefinisi sembarang
	   FS: Array bilangan bulat A terisi data sebanyak n elemen.
	       Elemen array terisi bilangan bulat menaik (ascending).
	       Jika bilangan yang dibaca lebih kecil dari
	       bilangan sebelumnya, data tidak masuk ke dalam array dan
	       pembacanaan terhenti.
	*/
	var x int
	fmt.Scan(&x)
	A[0] = x
	*n = 1
	for i := 1; i < NMAX; i++ {
		fmt.Scan(&x)
		if x < A[i-1] {
			break
		}
		A[i] = x
		*n++
	}
}

func cetakData(A tabInt, n int) {
	/*
		IS: Array bilangan A dan banyak elemen n terdefinisi
		FS: Tercetak di layar bilangan dengan format:
			"Data bilangan: <e1> <e2> <e3> ..<en>"
	*/
	fmt.Print("Data Bilangan: ")
	for i := 0; i < n; i++ {
		fmt.Print(A[i], " ")
	}
	fmt.Println()
}

func binarySearch(A tabInt, n int, x int, idx *int) {
	/*
		IS: Array bilangan A sebanyak n elemen dan bilangan x yang
		    dicari terdefinisi
		FS: idx berisi indeks lokasi x berada jika x ditemukan.
		    Jika tidak ditemukan bernilai -1.
	*/
	var kiri, kanan, tengah int
	kanan = n - 1
	tengah = (kiri + kanan) / 2
	*idx = -1
	for kiri < kanan && A[tengah] != x {
		if x < A[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
		tengah = (kiri + kanan) / 2
	}
	if A[tengah] == x {
		*idx = tengah
	}
}
