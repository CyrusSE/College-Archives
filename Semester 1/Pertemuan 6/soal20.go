package main

import "fmt"

func main() {
	var x, total int
	fmt.Print("Masukkan bilangan X : ")
	fmt.Scan(&x)
	for x > 0 {
		angka := x % 10
		fmt.Print(angka, " ")
		x = x / 10
		total = total + angka
	}
	fmt.Println()
	fmt.Println(total)
}
