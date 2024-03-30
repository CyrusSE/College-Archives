package main

import "fmt"

func main() {
	type mahasiswa struct {
		nama string
		nim  int
		ipk  float64
	}
	var orang1, orang2, orang3 mahasiswa
	fmt.Scan(&orang1.nama, &orang1.nim, &orang1.ipk)
	fmt.Scan(&orang2.nama, &orang2.nim, &orang2.ipk)
	fmt.Scan(&orang3.nama, &orang3.nim, &orang3.ipk)
	fmt.Printf("%.2f\n", (orang1.ipk+orang2.ipk+orang3.ipk)/3)
	if orang1.ipk > orang2.ipk && orang1.ipk > orang3.ipk {
		fmt.Println(orang1.nama)
	} else if orang2.ipk > orang3.ipk && orang2.ipk > orang1.ipk {
		fmt.Println(orang2.nama)
	} else {
		fmt.Println(orang3.nama)
	}
}
