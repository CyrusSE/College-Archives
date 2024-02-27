package main

import "fmt"

func main() {
	var x, y, z int
	var check bool
	fmt.Print("Masukkan bilangan ke-1 : ")
	fmt.Scan(&x)
	fmt.Print("Masukkan bilangan ke-2 : ")
	fmt.Scan(&y)
	fmt.Print("Masukkan bilangan ke-3 : ")
	fmt.Scan(&z)
	terkecil := 0
	terbesar := 0
	tengah := 0
	berapa := 0
	berapa2 := 0
	if x < y && x < z {
		terkecil = x
	} else if y < x && y < z {
		terkecil = y
	} else if z < x && z < y {
		terkecil = z
	}
	if x > y && x > z {
		terbesar = x
	} else if y > x && y > z {
		terbesar = y
	} else if z > x && z > y {
		terbesar = z
	}
	if x == terkecil && y == terbesar || x == terbesar && y == terkecil {
		tengah = z
	} else if y == terkecil && z == terbesar || y == terbesar && z == terkecil {
		tengah = x
	} else if x == terkecil && z == terbesar || x == terbesar && z == terkecil {
		tengah = y
	}
	for i := terkecil; i <= terbesar; i++ {
		berapa = berapa + 1
		if i == tengah {
			break
		}
	}
	for i := terbesar; i >= terkecil; i-- {
		berapa2 = berapa2 + 1
		if i == tengah {
			break
		}
	}
	if berapa == berapa2 {
		check = true
	}
	fmt.Print(check)
}
