package main

import "fmt"

func main() {
	var x int
	var sedia, kartu, diskon, cashback bool
	fmt.Print("Masukkan total belanja : ")
	fmt.Scan(&x)
	fmt.Print("Apakah bersedia untuk dibuatkan kartu? ")
	fmt.Scan(&sedia)
	diskon = (x >= 100000)
	kartu = (sedia && diskon)
	cashback = (x >= 200000 && sedia)
	// if x >= 100000 {
	// 	diskon = true
	// }
	// if sedia && diskon {
	// 	kartu = true
	// }
	// if x >= 200000 && sedia {
	// 	cashback = true
	// }
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
}
