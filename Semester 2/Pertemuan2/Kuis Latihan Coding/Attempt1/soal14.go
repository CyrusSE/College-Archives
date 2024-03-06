package main

import "fmt"

func nilaiMutlak(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	var x int
	fmt.Scan(&x)
	fmt.Println(nilaiMutlak(x))
}
