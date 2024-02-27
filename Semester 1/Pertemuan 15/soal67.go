package main

import "fmt"

func main() {
	var n, m int
	fmt.Print("Masukkan 2 bilangan positif : ")
	fmt.Scan(&n, &m)
	besar := 0
	if n > m {
		besar = n
	} else {
		besar = m
	}
	nomor := 0
	for i := 1; i <= besar; i++ {
		if n%i == 0 && m%i == 0 {
			nomor = i
		}
	}
	fmt.Print(nomor)
}
