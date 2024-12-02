package main

import "fmt"

func main() {
	var n, x, i, j, total int
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		for j = 1; j <= 7; j++ {
			fmt.Scan(&x)
			total += x
		}
		fmt.Println(total)
		total = 0
	}
}
