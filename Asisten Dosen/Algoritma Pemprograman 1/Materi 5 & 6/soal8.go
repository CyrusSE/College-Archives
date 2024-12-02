package main

import "fmt"

func main() {
	var n, total, i int
	total = 1
	fmt.Scan(&n)
	for i = n; i >= 1; i-- {
		total *= i
	}
	fmt.Println(total)
}
