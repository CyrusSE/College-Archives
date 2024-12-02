package main

import "fmt"

func main() {
	var n, i int
	var jam, total float64
	fmt.Scan(&n)
	for i = 1; i <= n; i++ {
		fmt.Scan(&jam)
		total += jam
	}
	fmt.Printf("%.3f", total/float64(n))
}
