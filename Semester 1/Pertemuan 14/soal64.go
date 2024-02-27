package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan bilangan x : ")
	fmt.Scan(&x)
	kecil := 1
	besar := x
	for i := 1; i <= x; i++ {
		for is := 1; is <= x; is++ {
			if i == kecil || i == besar {
				fmt.Print(i, " ")
			} else if is == kecil || is == besar {
				fmt.Print(i, " ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
