package main

import "fmt"

func main() {
	var x, y int
	var check bool
	fmt.Scan(&x, &y)
	check = y%x == 0
	fmt.Print(check)

}
