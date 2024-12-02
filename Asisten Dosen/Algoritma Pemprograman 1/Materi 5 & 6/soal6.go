package main

import "fmt"

func main() {
	var n, s, i int
	fmt.Scan(&n)
	for i = 0; i < n; i++ {
		fmt.Scan(&s)
		fmt.Println(s*s, 4*s)
	}
}
