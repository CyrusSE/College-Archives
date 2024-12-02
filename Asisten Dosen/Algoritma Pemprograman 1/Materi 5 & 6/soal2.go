package main

import "fmt"

func main() {
	var n, x, i int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&x)
		fmt.Println(x * 10)
	}
}
