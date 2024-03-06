package main

import "fmt"

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)
	fmt.Println(permutation(a, c), combination(a, c))
	fmt.Println(permutation(b, d), combination(b, d))
}
func factorial(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutation(n int, r int) int {
	return factorial(n) / factorial(n-r)
}

func combination(n int, r int) int {
	return factorial(n) / (factorial(r) * (factorial(n - r)))
}
