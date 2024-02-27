package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan bilangan x positif : ")
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		for i := 1; i <= x; i++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
