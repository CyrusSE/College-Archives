package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan nilai yang akan di konversi : ")
	fmt.Scan(&x)
	binary := ""
	konversi := 0
	for x > 0 {
		konversi = x % 2
		binary = fmt.Sprint(konversi, binary)
		x = x / 2
	}
	fmt.Println(binary)
}
