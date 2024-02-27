package main

import "fmt"

func main() {
	var x, y, hasil int
	fmt.Scan(&x)
	for x > 0 {
		hasil = x % 10
		y = y + hasil
		x = x / 10
	}
	fmt.Print(y)
}
