package main

import "fmt"

func main() {
	var x, biaya, biayakirim, totalbiaya, gram int
	fmt.Print("Masukkan berat parcel (gr) : ")
	fmt.Scan(&x)
	gram = x % 1000
	x = (x - gram) / 1000
	if x < 10 {
		if gram >= 500 {
			biaya = gram * 5
		} else {
			biaya = gram * 15
		}
	}
	biayakirim = x * 10000
	totalbiaya = biayakirim + biaya
	fmt.Println(x, "kg +", gram, "gr")
	fmt.Println("Rp.", biayakirim, "+", "Rp.", biaya, "=", totalbiaya)
}
