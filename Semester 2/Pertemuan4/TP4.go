package main

import "fmt"

func tribonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	if n == 3 {
		return 1
	}
	return tribonacci(n-1) + tribonacci(n-2) + tribonacci(n-3)
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(tribonacci(n))
}
