package main

import "fmt"

func main() {
	var n int
	var luas, keliling, jumlah int
	fmt.Print("Masukan n : ")
	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Print("Masukkan angka ke ", i, " : ")
		fmt.Scan(&jumlah)
		luas = jumlah * jumlah
		keliling = jumlah + jumlah + jumlah + jumlah
		fmt.Println(luas, keliling)
	}
}
