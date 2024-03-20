package main

import "fmt"

func cekGenap(x int) {
	/*
		I.S. terdefinisi bilangan bulat positif.
		F.S. tercetak "Ya" bila bilangan masukan adalah genap atau "Tidak" bila bukan.
	*/

	if x%2 == 0 {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	cekGenap(bilangan)
}
