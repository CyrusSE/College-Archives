package main

import "fmt"

func menu() {
	fmt.Println("-----------------------")
	fmt.Printf("%15s\n", "M E N U")
	fmt.Println("-----------------------")
	fmt.Println("1. Hitung Penjumlahan\n2. Hitung Perkalian\n3. Hitung Pembagian\n4. Exit\n-----------------------")
}

func hitungJumlah() {
	var x, y int
	fmt.Print("Masukkan dua bilangan yang akan dijumlahkan: ")
	fmt.Scan(&x, &y)
	fmt.Println("Hasil penjumlahan:", x+y)
}

func hitungKali() {
	var x, y int
	fmt.Print("Masukkan dua bilangan yang akan dikalikan: ")
	fmt.Scan(&x, &y)
	fmt.Println("Hasil perkalian:", x*y)
}

func hitungBagi() {
	var x, y float64
	fmt.Print("Masukkan dua bilangan yang akan dibagikan: ")
	fmt.Scan(&x, &y)
	fmt.Println("Hasil pembagian:", x/y)
}

func main() {
	var pilih int
	for pilih != 4 {
		menu()
		fmt.Print("Pilih (1/2/3/4)? ")
		fmt.Scan(&pilih)
		if pilih == 1 {
			hitungJumlah()
		} else if pilih == 2 {
			hitungKali()
		} else if pilih == 3 {
			hitungBagi()
		} else {
			break
		}
	}
}
