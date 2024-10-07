package main

import "fmt"

func main() {
	var r float64
	var luasL, kelilingL float64
	fmt.Scan(&r)
	luasL = 3.14 * r * r
	kelilingL = 2 * 3.14 * r
	fmt.Println(luasL)
	fmt.Println(kelilingL)
}
