package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan 1 bilangan positif : ")
	fmt.Scan(&x)
	if x%3 == 0 {
		fmt.Println("Kelipatan 3")
	}

	if x%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}
