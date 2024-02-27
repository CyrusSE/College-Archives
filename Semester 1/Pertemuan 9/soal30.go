package main

import "fmt"

func main() {
	var x int
	var check bool
	fmt.Scan(&x)
	check = (x%400 == 0 || x%4 == 0 && x%100 != 0)
	fmt.Print(check)
}
