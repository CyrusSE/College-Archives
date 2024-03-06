package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	for i := 1; i <= x; i++ {
		for b := 1; b <= x; b++ {
			if i == b {
				fmt.Print(" ")
			} else {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
