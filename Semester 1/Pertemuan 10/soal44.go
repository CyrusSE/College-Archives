package main

import "fmt"

func main() {
	var x int
	var check, diskon, kartu, cashback bool
	fmt.Print("Masukkan total belanjar dan bersedia/tidak : ")
	fmt.Scan(&x, &check)
	diskon1 := 0
	if x > 100000 {
		diskon1 = x * 10 / 100
		x = x - diskon1
		diskon = true
	}
	if check && diskon {
		kartu = true
	}
	if x > 200000 && kartu {
		cashback = true
		x = x - 75000
	}
	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println("Rp.", x)

}
