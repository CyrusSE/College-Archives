package main

import "fmt"

func gcd(a, b int) int {
	if b <= 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(gcd(A, B))
}
