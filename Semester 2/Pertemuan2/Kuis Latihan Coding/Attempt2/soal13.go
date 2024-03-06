package main

import "fmt"

func hitungDeret(a, b int) float64 {
	var hasil float64
	for i := a; i <= b; i++ {
		if (i-a)%2 == 0 {
			hasil += 3.0 / float64(i)
		} else {
			hasil -= 3.0 / float64(i)
		}
	}
	return hasil
}

func main() {
	var n, m int
	fmt.Scan(&n)
	fmt.Scan(&m)
	fmt.Printf("%.2f\n", hitungDeret(n, m))
}
