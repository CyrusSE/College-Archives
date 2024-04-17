package main

import "fmt"

const NMax int = 1000

type tabInt [NMax]int

func isiArray(arrInt *tabInt, n int) {
	/*I.S. Terdefinisi nilai bilangan bulat sebanyak n.
	  F.S. Array arrInt berisi data yang diberikan*/
	var bil, i int //variabel untuk input dan variabel untuk loop
	for i = 0; i < n; i++ {
		fmt.Scan(&bil)
		for bil < 0 { // loop untuk mengecek apakah inputan < 0 , jika benar maka meminta inputan ulang
			fmt.Scan(&bil)
		}
		arrInt[i] = bil
	}
}

func cekGanjil(x int) bool {
	/*Mengembalikan true apabila x adalah bilangan ganjil dan false jika sebaliknya*/
	return x%2 != 0
}

func prosesAngkaGanjil(arrInt tabInt, n int) {
	/*I.S. Terdefinisi array arrInt, array tidak kosong.
	  F.S. Menampilkan semua bilangan ganjil dalam array dan bilangan ganjil terbesar dalam array arrInt*/
	var i, tertinggi int //variabel untuk perulangan dan untuk menyimpan nilai tertinggi dari kumpulan angka ganjil
	tertinggi = 0
	for i = 0; i < n; i++ {
		if cekGanjil(arrInt[i]) {
			fmt.Print(arrInt[i], " ")  // menampilkan angka ganjil
			if arrInt[i] > tertinggi { // mencari nilai tertinggi dari kumpulan angka ganjil
				tertinggi = arrInt[i]
			}
		}
	}
	if tertinggi == 0 { // ketika tidak ada nilai ganjil tertinggi maka yang ditampilkan "-"
		fmt.Println("-")
	} else {
		fmt.Println()
	}
	fmt.Println(tertinggi)
	// tampilkan nilai tertinggi dari  kumpulan angka ganjil
}

func main() {
	var N int
	var arr tabInt
	fmt.Scan(&N)
	isiArray(&arr, N)
	prosesAngkaGanjil(arr, N)
}
