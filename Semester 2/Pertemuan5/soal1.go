package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n < 10 {
		n += 1
	} else {
		n += 2
	}
	fmt.Println(n)
}
