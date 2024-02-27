package main

import "fmt"

func main() {
	var x int
	var nilai bool
	fmt.Print("Masukkan bilang positif : ")
	fmt.Scanf("%d", &x)
	nilai = true
	for x > 0 {
		angka1 := x % 10
		angka2 := x % 100 / 10
		if angka2 != 0 {
			if angka1+angka2 != 1 && angka2-angka1 != 1 && angka1-angka2 != 1 {
				nilai = false
			}
		}
		x = x / 10
	}
	fmt.Print(nilai)
}
