package main

import "fmt"

func main() {
	var arr1 = [12]int{2, 3, 4, 9, 11, 13, 17, 18, 19, 20, 24, 26}
	kiri := 0
	kanan := 12
	tengah := (kiri + kanan) / 2
	cari := 26
	for kiri < kanan && arr1[tengah] != cari {
		if cari < arr1[tengah] {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
		//fmt.Println(tengah)
		tengah = (kiri + kanan) / 2
	}
	if arr1[tengah] != cari {
		fmt.Println("Gak nemu cuy")
	} else {
		fmt.Println("Nemu di index", tengah)
		fmt.Println(arr1[tengah])
	}
}
