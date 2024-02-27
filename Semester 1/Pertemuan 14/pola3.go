package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	for baris := 0; baris < n; baris++ {
		for kolom := 0; kolom < n; kolom++ {
			if baris == kolom || kolom == n-1-baris {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
