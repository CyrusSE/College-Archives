package main

import "fmt"

func main() {
	var x, d1, d2, d3 int
	fmt.Scan(&x)
	d1 = x / 100
	d2 = (x % 100) / 10
	d3 = x % 10
	fmt.Println(d1, d2, d3)
}
