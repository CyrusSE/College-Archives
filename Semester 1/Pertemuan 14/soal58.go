package main

import "fmt"

func main() {
	var x, check, terbesar int
	fmt.Scan(&x)
	for x > 0 {
		check = x % 10
		if check > terbesar {
			terbesar = check
		}
		x = x / 10
	}
	fmt.Print(terbesar)
}
