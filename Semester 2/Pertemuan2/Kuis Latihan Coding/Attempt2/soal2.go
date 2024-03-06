package main

import "fmt"

func luasPersegi(x int) int {
	return x * x
}

func main() {
	var sisi int
	fmt.Scan(&sisi)
	fmt.Println(luasPersegi(sisi))
}
