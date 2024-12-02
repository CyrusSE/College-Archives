package main

import "fmt"

func main() {
	var n, x, i int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&x)
		fmt.Print(x*i, " ")
	}
}
