package main

import "fmt"

func main() {
	var x, angka, angka2 int
	var selisih bool
	fmt.Print("Masukkan nilai x : ")
	fmt.Scan(&x)
	selisih = true
	for x > 0 && selisih {
		angka = x % 10
		x = x / 10
		angka2 = x % 10
		selisih = (angka-angka2 == 1 || angka2-angka == 1)
	}
	fmt.Print(selisih)
}
