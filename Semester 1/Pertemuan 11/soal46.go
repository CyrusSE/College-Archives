package main

import "fmt"

func main() {
	var umur int
	var KK bool
	fmt.Print("Masukkan usia : ")
	fmt.Scan(&umur)
	fmt.Print("Memiliki Kartu Keluarga atau tidak? ")
	fmt.Scan(&KK)
	if umur >= 17 && KK {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("belum bisa membuat KTP")
	}
}
