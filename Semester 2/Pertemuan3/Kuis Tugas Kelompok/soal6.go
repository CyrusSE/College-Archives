package main

import "fmt"

func persegi(a *int) {
	var l, k, n int
	for i := 1; i <= *a; i++ {
		fmt.Scan(&n)
		l = n * n
		k = 4 * n
		fmt.Println(l, k)
	}
}
func main() {
	var n int
	fmt.Scan(&n)
	persegi(&n)
}
