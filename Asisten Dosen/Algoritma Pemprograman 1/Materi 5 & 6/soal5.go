package main

import "fmt"

func main() {
	var n, i, total, sandi int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&sandi)
		total += (sandi % 10) + (sandi / 1000)
	}
	fmt.Println(total)
}
