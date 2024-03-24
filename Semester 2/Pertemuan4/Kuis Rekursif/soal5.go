package main

import "fmt"

func pola(n, start int) {
	if start == n {
		for i := 1; i <= start; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	} else {
		for i := 1; i <= start; i++ {
			fmt.Print("*")
		}
		fmt.Println()
		pola(n, start+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	pola(n, 1)
}
