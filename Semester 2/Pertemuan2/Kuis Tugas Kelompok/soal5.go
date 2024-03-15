package main

import "fmt"

func faktorial(n int) int {
	var hasil int
	hasil = 1
	for i := 1; i <= n; i++ {
		hasil = hasil * i
	}
	return hasil
}

func permutasi(x, y int) int {
	var hasil int
	if x >= y {
		hasil = faktorial(x) / faktorial(x-y)
	} else {
		hasil = faktorial(y) / faktorial(y-x)
	}
	return hasil
}

func main() {
	var x, y, hasil int
	fmt.Scan(&x, &y)
	hasil = permutasi(x, y)
	x = faktorial(x)
	y = faktorial(y)
	fmt.Println(x, y, hasil)
}
