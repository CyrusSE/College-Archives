package main

import "fmt"

func main() {
	var arr1 = [10]int{20, 19, 18, 17, 13, 11, 9, 4, 3, 2}

	kiri := 0
	kanan := 10
	tengah := (kiri + kanan) / 2
	cari := 18
	nemu := -1
	for kiri < kanan && nemu == -1 {
		if cari > arr1[tengah] {
			kanan = tengah - 1
		} else if cari < arr1[tengah] {
			kiri = tengah + 1
		} else {
			nemu = tengah
		}
		//fmt.Println(tengah)
		tengah = (kiri + kanan) / 2
	}
	if nemu == -1 {
		fmt.Println("Gak nemu cuy")
	} else {
		fmt.Println("Nemu di index", nemu)
		fmt.Println(arr1[nemu])
	}
}
