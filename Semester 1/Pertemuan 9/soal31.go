package main

import "fmt"

func main() {
	var x, y, z int
	var this bool
	fmt.Scan(&x, &y, &z)
	this = (x+y > z && x+z > y && y+z > x)
	fmt.Print(this)
}
