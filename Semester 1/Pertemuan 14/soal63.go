package main

import "fmt"

func main() {
	var x int
	fmt.Print("Masukkan bilangan x : ")
	fmt.Scan(&x)
	//check := 0
	for i := 1; i <= x; i++ {
		//check = i
		for is := 1; is <= x; is++ {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}
}
