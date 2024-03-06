package main

import (
	"fmt"
	"math"
)

func panjangX(R, T float64) float64 {
	return R * math.Cos(T*math.Pi/180)
}

func panjangY(R, T float64) float64 {
	return R * math.Sin(T*math.Pi/180)
}

func main() {
	var R, T float64
	fmt.Scan(&R, &T)
	x := panjangX(R, T)
	y := panjangY(R, T)
	fmt.Printf("%.2f %.2f", x, y)
}
