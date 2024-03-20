package main

import "fmt"

func koversiTemp(C float64, R, F, K *float64) {
	*F = 9.0/5.0*C + 32
	*R = 4.0 / 5.0 * C
	*K = C + 273.15
}

func main() {
	var C, R, F, K float64
	fmt.Scan(&C)
	koversiTemp(C, &R, &F, &K)
	fmt.Print(R, "R ", F, "F ", K, "K")
}
