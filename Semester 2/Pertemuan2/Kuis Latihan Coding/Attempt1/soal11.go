package main

import "fmt"

func sumRepeatUntillTen(x int) int {
	var hasil int
	for i := 1; i <= x*10; i++ {
		hasil += i
	}
	return hasil
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(sumRepeatUntillTen(N))
}
