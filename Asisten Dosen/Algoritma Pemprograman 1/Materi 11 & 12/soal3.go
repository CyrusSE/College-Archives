package main

import "fmt"

func main() {
	var x, total, angkaterakhir int
	fmt.Scan(&x)
	for x > 0 {
		angkaterakhir = x % 10
		fmt.Print(angkaterakhir, " ")
		total = total + angkaterakhir
		x = x / 10
	}
	fmt.Println()
	fmt.Println(total)
}
