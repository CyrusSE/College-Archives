package main

import "fmt"

func fibonacci(n, a, b, start int) {
	if start <= n {
		if start <= 1 {
			fmt.Println(start)
		} else {
			a, b = b, a+b
			fmt.Println(b)
		}
		fibonacci(n, a, b, start+1)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fibonacci(n, 0, 1, 0)
}
