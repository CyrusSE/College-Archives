package main

import "fmt"

func main() {
	var x, y, z int
	fmt.Print("Masukkan 3 bilangan positif : ")
	fmt.Scan(&x, &y, &z)
	if x == y && x == z {
		fmt.Print("Segitiga Sama Sisi")
	} else if x == z && x != y {
		fmt.Print("Segitiga Sama Kaki")
	} else {
		fmt.Print("Segitiga Sembarang")
	}
}
