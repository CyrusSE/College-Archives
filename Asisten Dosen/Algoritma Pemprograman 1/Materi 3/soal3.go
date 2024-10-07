package main

import "fmt"

func main() {
	var x, d1, d2, d3 int
	var terurut bool
	fmt.Scan(&x)
	d1 = x / 100
	d2 = (x % 100) / 10
	d3 = x % 10
	terurut = d1 > d2 && d2 > d3
	fmt.Println(terurut)
}
