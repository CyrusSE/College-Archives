package main

import "fmt"

func main() {
	var x, y int
	var check bool
	fmt.Print("Masukan angka untuk x : ")
	fmt.Scan(&x)
	fmt.Print("Masukan angka untuk y : ")
	fmt.Scan(&y)
	if (y % x) == 0 {
		check = true
	} else {
		check = false
	}
	fmt.Print(check)
}
