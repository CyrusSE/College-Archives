package main

import "fmt"

func jumlahGanjil(x int) int {
	var hasil int
	for i := 1; i <= x; i++ {
		if i%2 != 0 {
			hasil += i
		}
	}
	return hasil
}

func main() {
	var N int
	fmt.Scan(&N)
	fmt.Print(jumlahGanjil(N))
}
