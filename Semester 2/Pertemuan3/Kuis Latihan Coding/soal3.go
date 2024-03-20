package main

import "fmt"

func hitungKeliling(radius float64, kel *float64) {
	/* I.S.: Terdefinisi radius lingkaran
	   F.S.: kel berisi hasil perhitungan keliling lingkaran */

}

func main() {
	var radius, keliling float64
	fmt.Scan(&radius)
	hitungKeliling(radius, &keliling)
	fmt.Println(keliling)
}
