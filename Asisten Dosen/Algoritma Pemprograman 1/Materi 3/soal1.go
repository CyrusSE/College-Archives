package main

import "fmt"

func main() {
	var x int
	var check bool
	fmt.Scan(&x)
	check = x%2 != 0
	fmt.Println(check)
}
