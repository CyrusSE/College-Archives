package main

import "fmt"

func fibonacci(N int) int {
	var a, b, hasil, total int
	b = 1
	for i := 0; i < N; i++ {
		if i <= 1 {
			hasil = i
		} else {
			hasil = a + b
			a = b
			b = hasil
		}
		total += hasil
	}
	return total
}

func main() {
	var N, hasil int
	fmt.Scan(&N)
	hasil = fibonacci(N)
	fmt.Print(hasil)
}
