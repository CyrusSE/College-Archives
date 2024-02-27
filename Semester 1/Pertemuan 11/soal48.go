package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Masukkan 3 bilangan bulat : ")
	fmt.Scan(&a, &b, &c)
	terkecil := 0
	tengah := 0
	terbesar := 0
	if a >= b && a >= c {
		terbesar = a
	} else if b >= a && b >= c {
		terbesar = b
	} else if c >= a && c >= b {
		terbesar = c
	}
	if a <= b && a <= c {
		terkecil = a
	} else if b <= a && b <= c {
		terkecil = b
	} else if c <= a && c <= b {
		terkecil = c
	}
	if a >= b && a <= c || a <= b && a >= c {
		tengah = a
	} else if b >= a && b <= c || b <= a && b >= c {
		tengah = b
	} else if c >= a && c <= b || c <= a && c >= b {
		tengah = c
	}
	fmt.Println(terkecil, tengah, terbesar)
}
